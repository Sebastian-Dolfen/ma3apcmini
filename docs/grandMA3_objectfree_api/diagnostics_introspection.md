# Diagnostics & introspection

Dump the API descriptor, read text-screen lines, and sample frame-rate / per-frame timing for debugging.

## Functions

- [`GetApiDescriptor`](#getapidescriptor)
- [`GetObjApiDescriptor`](#getobjapidescriptor)
- [`GetTextScreenLine`](#gettextscreenline)
- [`GetTextScreenLineCount`](#gettextscreenlinecount)
- [`GetDebugFPS`](#getdebugfps)
- [`GetSample`](#getsample)

---


## GetApiDescriptor

**Signature**

```
GetApiDescriptor(nothing): table of {string:function_name, string:arguments, string:return_values}
```

**Help page title:** `GetApiDescriptor()`

**Description**

The **GetApiDescriptor** Lua function returns a table with a description of all the object-free Lua functions. These are descriptions only. The function does not have any actual functions. The table is not sorted.

**Arguments**

This function does not accept any arguments.

**Return**

- **Table**:

            The returned table contains elements with other tables.

  - **String**:

                This is the Api function name.

  - **String**:

                This is the description of the Api arguments.

  - **String**:

                This is the description of the Api returns.

**Example**

This example prints the content of the returned table.

```lua
return function ()
    -- This returns information about all the Lua "object-free" functions.
    -- GetApiDescriptor() returns a table with all the functions. 
    -- Each table element is another table with the name, argument description, and return description.
      for key,value in ipairs(GetApiDescriptor()) do
        if value[1] ~= nil then
          Printf("Api " .. key .. " is: " .. value[1])
        end
        if value[2] ~= nil then
          Printf("Arguments: " .. value[2])
        end
        if value[3] ~= nil then
          Printf("Returns: " .. value[3])
        end
        Printf("---------------------------------------")
      end
  end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getapidescriptor.html>


## GetObjApiDescriptor

**Signature**

```
GetObjApiDescriptor(nothing): table of {string:function_name, string:arguments, string:return_values}
```

**Help page title:** `GetObjApiDescriptor()`

**Description**

The **GetObjApiDescriptor** Lua function returns a table with a description of all the object Lua functions. These are descriptions only. The function does not have any actual functions. The table is not sorted.

**Arguments**

This function does not accept any arguments.

**Return**

- **Table**:

 The returned table contains elements with three values.

  - **String**:

 This is the API function name.

  - **String**:

 This is the description of the API arguments.

  - **String**:

 This is the description of the API returns.

**Example**

This example prints the content of the returned table.

```lua
return function ()
  -- This returns information about all the Lua "object" functions.
  -- GetObjApiDescriptor() returns a table with all the functions. 
  -- Each table element is another table with the name, argument description, and return description.
    for key,value in ipairs(GetObjApiDescriptor()) do
      if value[1] ~= nil then
        Printf("Api " .. key .. " is: " .. value[1])
      end
      if value[2] ~= nil then
        Printf("Arguments: " .. value[2])
      end
      if value[3] ~= nil then
        Printf("Returns: " .. value[3])
      end
      Printf("---------------------------------------")
      end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getobjapidescriptor.html>


## GetTextScreenLine

**Signature**

```
GetTextScreenLine(nothing): integer:internal line number
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetTextScreenLineCount

**Signature**

```
GetTextScreenLineCount([integer:starting internal line number]): integer:line count
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetDebugFPS

**Signature**

```
GetDebugFPS(nothing): float:fps
```

**Help page title:** `GetDebugFPS()`

**Description**

The **GetDebugFPS** Lua function returns a float number with the frames per second.

**Arguments**

This function does not accept any arguments.

**Return**

- **Number**:

 The returned number indicates the current frames per second.

**Example**

This example prints the FPS number:

```lua
return function ()
    -- Prints the current frames per second.
    Printf("Current FPS: " .. GetDebugFPS())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getdebugfps.html>


## GetSample

**Signature**

```
GetSample(string:type('MEMORY', 'CPU', 'CPUTEMP', 'GPUTEMP', 'SYSTEMP', 'FANRPM')): integer:current_value_in_percent
```

**Help page title:** `GetSample(string)`

**Description**

The **GetSample** Lua function returns a number representing a percentage usage based on a string input.

**Arguments**

- **String**:

 Only a specific list of strings can be input:

  - MEMORY

  - CPU

  - CPUTEMP

  - GPUTEMP

  - SYSTEMP

  - FANRPM

**Return**

- **Number**:

 A number (float) is returned.

**Example**

This example stores the different samples in a table and then prints the content of the table:

```lua
return function()
    -- Gather the sample information in a table
    local sample = {}
    sample["MEMORY"] = GetSample("MEMORY")
    sample["CPU"] = GetSample("CPU")
    sample["CPUTEMP"] = GetSample("CPUTEMP")
    sample["GPUTEMP"] = GetSample("GPUTEMP")
    sample["SYSTEMP"] = GetSample("SYSTEMP")
    sample["FANRPM"] = GetSample("FANRPM")
    -- Print the collected data
    Printf("Memory ; ".. sample["MEMORY"])
    Printf("CPU ; ".. sample["CPU"])
    Printf("CPU temperature ; ".. sample["CPUTEMP"])
    Printf("GPU temperature ; ".. sample["GPUTEMP"])
    Printf("System temperature ; ".. sample["SYSTEMP"])
    Printf("Fan RPM ; ".. sample["FANRPM"])
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getsample.html>

---

[Back to index](README.md)
