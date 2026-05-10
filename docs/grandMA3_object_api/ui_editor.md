# UI Editor / Settings

Looking up the UI editor and UI settings windows associated with an object.

## Contents

- [GetUIEditor](#getuieditor)
- [GetUISettings](#getuisettings)

---

## GetUIEditor

**Signature:**

```
GetUIEditor(light_userdata:handle): string:ui_editor_name
```

**Summary:** The GetUIEditor function returns a text string with the name of the UI editor for the object.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.

**Returns:**

- String : The function returns a text string with the name of the object's UI editor.

**Example:**

This example prints the name of the selected sequence's editor.

```lua
return function()
    -- SelectedSequence() creates a handle to the selected sequence.
    local selectedSequence = SelectedSequence()
    -- Get the name of the editor for the sequence object.
    local seqEditor = selectedSequence:GetUIEditor()
    -- Print some feedback.
    if seqEditor ~= nil then
        Printf("The name of the editor is: " .. seqEditor)
    else
        Printf("The object doesn not appear to have an editor.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_getuieditor.html>

---

## GetUISettings

**Signature:**

```
GetUISettings(light_userdata:handle): string:ui_settings_name
```

**Summary:** The GetUISettings function returns a text string with the name of the UI settings for the object.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.

**Returns:**

- String : The function returns a text string with the name of the object's UI settings.

**Example:**

This example prints the name of the selected sequence's settings.

```lua
return function()
    -- SelectedSequence() creates a handle to the selected sequence.
    local selectedSequence = SelectedSequence()
    -- Get the name of the editor for the sequence object.
    local seqSettings = selectedSequence:GetUISettings()
    -- Print some feedback.
    if seqSettings ~= nil then
        Printf("The name of the settings is: " .. seqSettings)
    else
        Printf("The object doesn not appear to have an editor.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_getuisettings.html>

---

[Back to index](README.md)
