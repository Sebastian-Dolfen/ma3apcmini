# Import / Export

Reading and writing objects to XML files on disk.

## Contents

- [Export](#export)
- [Import](#import)

---

## Export

**Signature:**

```
Export(light_userdata:handle, string:file_path, string:file_name): boolean:success
```

**Summary:** The Export object Lua function exports an object into an XML file.

**Arguments:**

1. Handle : The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.
2. String : This is a string with the file path for the exported file.
3. String : This is a string containing the file name of the exported file.

**Returns:**

- Boolean : The function returns a boolean indicating if the export was a success.

**Used by APC plugin for:** saving plugin state or selected objects to disk.

**Example:**

This example exports the selected sequence into an XML file:

```lua
return function()
    --SelectedSequence() creates a handle to the selected sequence.
    local selectedSequence = SelectedSequence()
    if selectedSequence == nil then
        ErrPrintf("The selected sequence could not be found.")
        return
    end
    --The path is stored in a variable.
    local exportPath = GetPath(Enums.PathType.UserSequences)
    --The actual export function.
    local success = selectedSequence:Export(exportPath, "mySelectedSequence.xml")
    --Print some feedback.
    if success then
        Printf("The sequence is exported to: " .. exportPath)
    else
        ErrPrintf("The sequence could not be exported.")
    end
end
```

**Related:**

- Import - object function used to import an XML table.

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_export.html>

---

## Import

**Signature:**

```
Import(light_userdata:handle, string:file_path, string:file_name): boolean:success
```

**Summary:** The Import object Lua function imports an object written in XML format.

Restriction:
The imported files need to exist already to be imported.
Important:
The Lua import will merge the content of the XML file into the object without any confirmation pop-up.

**Arguments:**

1. Handle : The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.
2. String : This is a string with the path to the file location.
3. String : This is a string containing the file name of the desired file.

**Returns:**

- Boolean : The function returns a boolean indicating if the import was a success.

**Used by APC plugin for:** loading saved show fragments from disk for the bridge.

**Example:**

This example imports the content of an XML file into the selected sequence. The file is called "MySelectedSequence", and it is located at "../gma3_library/datapools/sequences". The file can be created using the example in the Export object function .

```lua
return function()
    --SelectedSequence() creates a handle to the selected sequence.
    -- The imported object will be merged into this sequence.
    local selectedSequence = SelectedSequence()
    -- Check if there is a selected sequence - if not then exit the function.
    if selectedSequence == nil then
        ErrPrintf("The selected sequence could not be found.")
        return
    end
    --The path is stored in a variable.
    local path = GetPath(Enums.PathType.UserSequences)
    --The actual import function.
    local success = selectedSequence:Import(path, "mySelectedSequence.xml")
    --Print some feedback.
    if success then
        Printf("The sequence is imported from: " .. path .. GetPathSeparator() .. "mySelectedSequence.xml")
    else
        ErrPrintf("The object could not be imported.")
    end
end
```

**Related:**

- Export - object function used to export an XML table.

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_import.html>

---

[Back to index](README.md)
