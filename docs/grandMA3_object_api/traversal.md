# Traversal

Walking the object tree: enumerating children, counting them, and fetching one by index.

## Contents

- [Children](#children)
- [Count](#count)
- [Ptr](#ptr)

---

## Children

**Signature:**

```
Children(light_userdata:handle): {light_userdata:child_handles}
```

**Summary:** The Children Lua function creates a table of handles for the children of an object.

**Arguments:**

1. Handle : The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object

**Returns:**

- Table : The function returns a table with handles for the child objects. If there are no children, then it returns an empty table.

**Used by APC plugin for:** iterating cues, executors, sequences, and data-pool entries when scanning what an APC button maps to.

**Example:**

This example returns the name of the cues in the first sequence of the selected data pool:

```lua
return function()
    -- Stores the handle for sequence 1 in a variable.
    local mySequence = DataPool().Sequences[1]
    if mySequence ~= nil then
            -- Use the "Children()" funciton to store a table with all the children in a new variable.
        local cues = mySequence:Children()
        -- For loop that uses the length operator on the cue variable.
        for i = 1, #cues do
            -- Text is printed for each child.
            Printf("Sequence 1 Child " .. i .. " = " .. cues[i].name)
        end
    else
        ErrPrintf("Sequence could not be found.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_children.html>

---

## Count

**Signature:**

```
Count(light_userdata:handle): integer:child_count
```

**Summary:** The Count function returns an integer number indicating the number of child objects.

**Arguments:**

1. Handle: The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.

**Returns:**

- Integer : The function returns an integer indicating the number of children of the object.

**Used by APC plugin for:** sizing the number of children before iterating (e.g., total cues in a sequence).

**Example:**

This example prints the selected sequence's number of children (cues).

```lua
return function()
    local numberChildren = SelectedSequence():Count()
    Printf("The selected Sequence has " .. numberChildren .. " cues.")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_count.html>

---

## Ptr

**Signature:**

```
Ptr(light_userdata:handle, integer:index(1-based)): light_userdata:child_handle
```

**Summary:** The Ptr Lua function returns the handle to a child object.

**Arguments:**

1. Handle : The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.
2. Integer : This is the index number for the desired child object. This index is 1-based.

**Returns:**

- Handle | nil : The function returns a handle to the child object. If the child object does not exist, then it returns nil.

**Used by APC plugin for:** random-access fetch of a child handle by index when binding a button to executor N.

**Example:**

This example prints the data connected to the first child of the selected sequence. It uses the Dump() function.

```lua
return function()
    -- SelectedSequence() creates a handle to the selected sequence.
    local selectedSequence = SelectedSequence()
    -- Check that a handle was returned - if not then exit function.
    if selectedSequence == nil then
        ErrPrintf("There is no selected sequence.")
        return
    end
    -- Get a handle to the first child object.
    local firstChild = selectedSequence:Ptr(1)
    -- Print some feedback.
    if firstChild ~= nil then
        Printf("=============== START OF DUMP ===============")
        firstChild:Dump()
        Printf("================ END OF DUMP ================")
    else
        ErrPrintf("The object do not have a child object.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_ptr.html>

---

[Back to index](README.md)
