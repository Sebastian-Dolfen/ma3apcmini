# Properties

Reading object property values and inspecting full object state.

## Contents

- [Get](#get)
- [Dump](#dump)

---

## Get

**Signature:**

```
Get(light_userdata:handle, string:property_name[, integer:role(Enums.Roles)]): light:userdata:child or string:property (if 'role' provided - always string)
```

**Summary:** The Get function returns a string with information about a specified property of the object, for instance, the object's name, class, or path.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.
2. String : The string must be the name of a valid property for the object.
3. Integer (optional): A valid role integer can be supplied. This will make the returned value a text string.

**Returns:**

- String: The function returns the value of the property. If the property is a boolean, then the return is "0" or "1" unless a role is defined (see optional integer argument above). When the role is supplied, a boolean is returned as "No" or "Yes".

**Used by APC plugin for:** reading object properties (e.g., No, Name, Appearance) to drive LED color/state on the APC.

**Example:**

This example prints information about the "Tracking" property of the selected sequence.

```lua
return function ()
    -- SelectedSequence() creates a handle to the selected sequence.
    local selectedSequence = SelectedSequence()
    -- Check if there is a selected sequence. If not, then exit the function.
    if selectedSequence == nil then
        ErrPrintf("The selected sequence could not be found.")
        return
    end
    -- Set a variable with the property name.
    local propertyName = "Tracking"
    -- Get the value of the property.
    local propertyValue = selectedSequence:Get(propertyName)
    local propertyValueString = selectedSequence:Get(propertyName, Enums.Roles.Edit)
    -- Return some feedback.
    if propertyValue ~= nil then
        Printf("The selected sequence's property " .. propertyName.. " has the value '" .. propertyValue .. "' and a string value of '" .. propertyValueString .. "'.")
    else
        ErrPrintf("The property could not be found.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_get.html>

---

## Dump

**Signature:**

```
Dump(light_userdata:handle): string:information
```

**Summary:** The Dump function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation for object-oriented calls.

**Returns:**

- The function returns nothing but outputs information about the object in the Command Line History window .

**Used by APC plugin for:** debug logging of object state in plugin development.

**Example:**

These examples all print information about the selected sequence in the Command Line History.
The first example using the colon operator:

```lua
return function ()
    -- Dump() is called on a function
    Printf("=============== START OF DUMP ===============")
    SelectedSequence():Dump()
    Printf("================ END OF DUMP ================")
end
```

The second example uses a variable with the same result:

```lua
return function ()
    --Stores the handle for the selected sequence in a local variable.
    local mySeqHandle = SelectedSequence()
    -- Dump() is called on the variable.
    Printf("=============== START OF DUMP ===============")
    mySeqHandle:Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_dump.html>

---

[Back to index](README.md)
