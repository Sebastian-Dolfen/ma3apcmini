# Playback State and Relationships

Querying runtime playback status and the dependency / reference graph of an object.

## Contents

- [HasActivePlayback](#hasactiveplayback)
- [GetReferences](#getreferences)
- [GetDependencies](#getdependencies)

---

## HasActivePlayback

**Signature:**

```
HasActivePlayback(light_userdata:handle): boolean:result
```

**Summary:** The HasActivePlayback Lua function returns a boolean indicating if an object has a currently active playback, for instance, if a sequence has an active cue.

**Arguments:**

1. Handle : The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.

**Returns:**

- Boolean : The function returns a boolean indicating the playback status: True : There is active playback. False : There is no active playback.

**Used by APC plugin for:** lighting the executor LED red/green based on whether the sequence is currently playing.

**Example:**

To return the information if the selected sequence has an active playback, create a plugin with this code:

```lua
return function()
    -- Stores the handle of the selected sequence.
    local selectedSequence = SelectedSequence()

    -- The following 'if' gives different feedback based on the playback status.
    if selectedSequence:HasActivePlayback() then
        Printf("Sequence '" ..selectedSequence.name.. "' has active playback.")
    else
        Printf("Sequence '" ..selectedSequence.name.. "' has NO active playback.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_hasactiveplayback.html>

---

## GetReferences

**Signature:**

```
GetReferences(light_userdata:handle): {light_userdata:handle}
```

**Summary:** The GetReferences function returns a table with handles for the objects referencing this object.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.

**Returns:**

- Table : The function returns a table with the handles for the different objects referencing this object.

**Example:**

This example prints a dump of the selected sequence's first object in the returned table.

```lua
return function()
    -- SelectedSequence() creates a handle to the selected sequence.
    local selectedSequence = SelectedSequence()
    -- Get the references for the sequence.
    local seqReferences = selectedSequence:GetReferences()
    -- Check if there are any references and output a relevant feedback.
    if seqReferences ~= nil then
        -- There is a reference table returned. Print a dump of the first table element.
        Printf("=============== START OF DUMP ===============")
        seqReferences[1]:Dump()
        Printf("================ END OF DUMP ================")
    else
        Printf("No references found")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_getreferences.html>

---

## GetDependencies

**Signature:**

```
GetDependencies(light_userdata:handle): {light_userdata:handle}
```

**Summary:** The GetDependencies function returns a table with the objects' dependencies.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.

**Returns:**

- Table : The function returns a table with the handles for the different dependency objects.

**Example:**

This example prints a dump of the selected sequence's first object in the returned table.

```lua
return function()
    -- SelectedSequence() creates a handle to the selected sequence.
    local selectedSequence = SelectedSequence()
    -- Get the dependcies for the sequence.
    local seqDependencies = selectedSequence:GetDependencies()
    -- Check if there are any dependicies and output a relevant feedback.
    if seqDependencies ~= nil then
        -- There is a dependency table returned. Print a dump of the first table element.
        Printf("=============== START OF DUMP ===============")
        seqDependencies[1]:Dump()
        Printf("================ END OF DUMP ================")
    else
        Printf("No dependencies found")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_getdependencies.html>

---

[Back to index](README.md)
