# Addressing

Converting object handles into address strings for use in command-line operations.

## Contents

- [Addr](#addr)
- [AddrNative](#addrnative)
- [ToAddr](#toaddr)

---

## Addr

**Signature:**

```
Addr(light_userdata:handle[, light_userdata:base_handle[, boolean:force_parent-based_address[, boolean:force_commandline_index-based_address]]]): string:numeric_root_address
```

**Summary:** The Addr Lua object function converts a handle to an address string that can be used in commands.

See the Handle topic for more info regarding handles and links to other related functions.

**Arguments:**

1. Handle : The function takes a handle of the type "light_userdata" as an argument. This is the handle to the object where the address is requested. It can be omitted when using the colon notation on an object.
2. Handle (optional): The returned address is from the root as a default. This optional handle can specify a different base location. It must still be a base location in the address path from the root to the object.
3. Boolean | nil (optional): This can be useful if there is a difference between the ToAddr() and Addr(). Setting this to "true" uses the index number from the ToAddr() instead of the Addr() index number. See the example below.
4. Boolean (optional): In some edge cases, the cue address is not resolved correctly. Setting this boolean to true will fix this.

**Returns:**

- String : Text string with the address in a parent-child number format separated by dots.

**Used by APC plugin for:** persisting a stable text reference to an object across show reloads.

**Example:**

This example prints different versions of the address to a cue in a sequence:

```lua
return function()
    -- Creates a cue in sequence 1
    Cmd("Store Sequence 1 Cue 100 /Merge /NoConfirmation")
    --Store a handle to the created cue
    local cueObject = ObjectList("Sequence 1 Cue 100")[1]
    --Print different version of the handle address
    Printf("ToAddr:              " .. cueObject:ToAddr())
    Printf("Addr:                " .. cueObject:Addr())
    Printf("Addr(Parent, false, false): " .. cueObject:Addr(cueObject:Parent(), false, false))
    Printf("Addr(Parent, true, false):  " .. cueObject:Addr(cueObject:Parent(), true, false))
    Printf("Addr(Parent, false, true): " .. cueObject:Addr(cueObject:Parent(), false, true))
    Printf("Addr(Parent, true, true):  " .. cueObject:Addr(cueObject:Parent(), true, true))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_addr.html>

---

## AddrNative

**Signature:**

```
AddrNative(light_userdata:handle, light_userdata:base_handle[, boolean:escape_names]]): string:numeric_root_address
```

**Summary:** The AddrNative Lua object function converts a handle to an address string that can be used in commands.

See the Handle topic for more info regarding handles and links to other related functions.

**Arguments:**

1. Handle : The function takes a handle as an argument. This is the handle to the object where the address is requested. It can be omitted when using the colon notation on an object.
2. Handle (optional): The returned address is from the root as a default. This optional handle can specify a different base location. It still needs to be a base location in the address path from the root to the object.
3. Boolean (optional): Set this to "true" to get the returned names in quotation marks.

**Returns:**

- String : Text string with the address in a parent-child name format separated by dots.

**Used by APC plugin for:** building command-line addresses for Cmd() calls.

**Example:**

This example prints the address of the first sequence:

```lua
return function()
    -- Stores the handle to the first sequence.
    local mySequence = DataPool().Sequences[1]
    -- Print the native address.
    Printf("The full address is: " .. mySequence:AddrNative())
    -- Stores a handle to the default DataPool.
    local myDataPool = DataPool()
    -- Print the native address to the datapool using the default datapool as a base.
    Printf("The address in the datapool is: " .. mySequence:AddrNative(myDataPool))
    -- Print the native address to the datapool, using the default datapool as a base, with names as strings.
    Printf("The address in the datapool with quotes around the names is: " .. mySequence:AddrNative(myDataPool, true))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_addrnative.html>

---

## ToAddr

**Signature:**

```
ToAddr(light_userdata:handle,boolean:with_name[, boolean:use_visible_addr]): string:address
```

**Summary:** The ToAddr Lua object function converts a handle to an address string that can be used in commands.

See the Handle topic for more info regarding handles and links to other related functions.

**Arguments:**

1. Handle : The function takes a handle of the type "light_userdata" as an argument. It can be omitted when using the colon notation on an object.
2. Boolean : Set this to "true" to get the returned name. "False" will return the object type and index number.

**Returns:**

- String : Text string with the address.

**Used by APC plugin for:** human-readable address strings for log lines and command building.

**Example:**

This example returns the address of the first sequence of the selected data pool, prints the address in the Command Line History, and creates a grandMA3 command with a "Go" keyword in front of the address. This command is sent to the grandMA3 command line.

```lua
return function()
    -- Stores the handle in a variable.
    local mySequence = DataPool().Sequences[1]
    if mySequence ~= nil then
        -- Converts the handle to the address and store in variable.
        local mySequenceAddressName = mySequence:ToAddr(true)
        local mySequenceAddress = mySequence:ToAddr(false)
        -- Print the address to the Command Line History.
        Printf("The named address of the sequence is: " .. mySequenceAddressName)
        Printf("The system address of the sequence is: " .. mySequenceAddress)
        -- Send a 'Go' command with the address appended.
        Cmd("Go %s", mySequenceAddress)
    else
        ErrPrintf("The sequence could not be found")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_toaddr.html>

---

[Back to index](README.md)
