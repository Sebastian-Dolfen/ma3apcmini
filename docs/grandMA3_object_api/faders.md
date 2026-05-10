# Faders

Reading and writing fader values on objects that have them (sequences, executors, etc.). The central area of interest for the APCmini bridge.

## Contents

- [GetFader](#getfader)
- [GetFaderText](#getfadertext)
- [SetFader](#setfader)

---

## GetFader

**Signature:**

```
GetFader(light_userdata:handle, {[string:token(Fader*)], [integer:index]}): float:value[0..100]
```

**Summary:** The GetFader function returns a float number indicating a fader position for the object.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object. See the examples below.
2. Table : The table can contain two different elements: Token and Index. The important element is the token. Token : This is used to specify which fader the value is requested for. These are the valid values: FaderMaster FaderX FaderXA FaderXB FaderTemp FaderRate FaderSpeed FaderHighlight FaderLowlight FaderSolo FaderTime

**Returns:**

- Number : The function returns a float number indicating the fader position.

**Used by APC plugin for:** reading current executor fader position to send back to the APC for visual feedback.

**Example:**

This example prints the fader positions of the Master and Rate faders for the selected sequence.

```lua
return function()
    -- SelectedSequence() creates a handle to the selected sequence.
    local selectedSequence = SelectedSequence()
    -- Get the value for the Master fader.
    local faderMasterValue = selectedSequence:GetFader({})
    -- Get the value for the Rate fader.
    local faderRateValue = selectedSequence:GetFader({token="FaderRate"})
    -- Print feedback with the values.
    Printf("The selected sequence Master fader value is: ".. tostring(faderMasterValue))
    Printf("The selected sequence Rate fader value is: ".. tostring(faderRateValue))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_getfader.html>

---

## GetFaderText

**Signature:**

```
GetFaderText(light_userdata:handle, {[string:token(Fader*)], [integer:index]}): string:text
```

**Summary:** The GetFaderText function returns a text string indicating a fader value for the object.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object. See the examples below.
2. Table : The table can contain two different elements: Token and Index. The important element is the token. Token : This is used to specify which fader the value is requested for. These are the valid values: FaderMaster FaderX FaderXA FaderXB FaderTemp FaderRate FaderSpeed FaderHighlight FaderLowlight FaderSolo FaderTime

**Returns:**

- String : The function returns a text string indicating the fader value.

**Used by APC plugin for:** showing fader value labels in plugin UI.

**Example:**

This example prints the fader value text of the Master and Rate faders for the selected sequence.

```lua
return function()
    -- SelectedSequence() creates a handle to the selected sequence.
    local selectedSequence = SelectedSequence()
    -- Get the value for the Master fader. Since it is the default, no token needs to be defined.
    local faderMasterText = selectedSequence:GetFaderText({})
    -- Get the value for the Rate fader.
    local faderRateText = selectedSequence:GetFaderText({token="FaderRate"})
    -- Print feedback with the values.
    Printf("The selected sequence Master fader value text is: ".. tostring(faderMasterText))
    Printf("The selected sequence Rate fader value text is: ".. tostring(faderRateText))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_getfadertext.html>

---

## SetFader

**Signature:**

```
SetFader(light_userdata:handle, {[float:value[0..100]], [boolean:faderEnabled], [string:token(Fader*)]}): nothing
```

**Summary:** The Set Fader function sets a fader to a specified level. It must be used on an object that has faders.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.
2. Table : The table can contain up to three named elements using the key/value methods. "value" : This is a float number indicating the fader position on a scale from 0 to 100. This should always be part of the table. "token" : This is a string indicating the fader. The string must start with "Fader". It can be omitted, and then the value will be assigned to the Master fader. The fader name must be valid for the object being used. Possible tokens include: "FaderMaster" "FaderX" "FaderXA" "FaderXB" "FaderTemp" "FaderRate" "FaderSpeed" "FaderHighlight" "FaderLowlight" "FaderTime" "FaderSolo" "faderEnabled" : If the fader can be toggled, then this boolean can be used to enable or disable the fader. A true value sets the fader to enabled.

**Returns:**

- This function does not return anything.

**Used by APC plugin for:** driving an executor master from an APC fader (this is the central use case for the bridge).

**Example:**

This example changes the selected sequences' Master fader to 100% and the Time fader to 5 seconds and enables the time fader.

```lua
return function()
    -- SelectedSequence() creates a handle to the selected sequence.
    local selectedSequence = SelectedSequence()
    -- Set the master fader to 100. The FaderMaster is the default token, so it can be omitted.
    selectedSequence:SetFader({value=100.0})
    -- Set the time fader to 5 seconds and enable the fader.
    selectedSequence:SetFader({value=50.0, faderEnabled=1, token="FaderTime"})
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_setfader.html>

---

[Back to index](README.md)
