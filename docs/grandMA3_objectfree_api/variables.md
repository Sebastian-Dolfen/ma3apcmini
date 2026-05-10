# Variables (Global/User/Plugin/Addon)

Access and mutate the four variable scopes (global / per-user / per-plugin / per-addon).

## Functions

- [`GlobalVars`](#globalvars) [APC]
- [`UserVars`](#uservars) [APC]
- [`PluginVars`](#pluginvars) [APC]
- [`AddonVars`](#addonvars)
- [`SetVar`](#setvar) [APC]
- [`GetVar`](#getvar) [APC]
- [`GetVarVersion`](#getvarversion)
- [`DelVar`](#delvar)

---


## GlobalVars

**Signature**

```
GlobalVars(nothing): light_userdata:global_variables
```

**Help page title:** `GlobalVars()`

**Used by APC plugin for:** shared state across all users (e.g. last-touched executor).

**Description**

The GlobalVars function returns a handle to the set of global variables. Read more about these in the Variables topic in the Macro section.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The function returns a handle of the set of global variables.

**Example**

This example sets, gets, and deletes a global variable:

```lua
return function()
    -- Stores a local Lua variable with the handle for the global variable set.
    local variableSet = GlobalVars()
    -- Sets a global variable with an integer value using the SetVar() function.
    SetVar(variableSet, "myGlobalVar", 42)
    -- Prints the global variable using the GetVar() function.
    Printf("The value  of myGlobalVar is: " .. GetVar(variableSet, "myGlobalVar"))
    -- Deletes the global variable using the DelVar() function.
    DelVar(variableSet, "myGlobalVar")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_globalvars.html>


## UserVars

**Signature**

```
UserVars(nothing): light_userdata:user_variables
```

**Help page title:** `UserVars()`

**Used by APC plugin for:** per-user state (current page selection on the APC).

**Description**

The UserVars function returns a handle to the set of user variables. Read more about these in the Variables topic in the Macro section.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

            The function returns a handle of the set of user variables.

**Example**

This example sets, gets, and deletes a user variable:

```lua
return function()
    -- Stores a local Lua variable with the handle for the user variables.
    local variableSection = UserVars()

    -- Sets a user variable with an integer value using the SetVar function.
    SetVar(variableSection, "myUserVar", 42)

    -- Prints the user variable using the GetVar function.
    Printf("The value  of myUserVar is: " .. GetVar(variableSection, "myUserVar"))

    -- Deletes the user variable using the DelVar function.
    DelVar(variableSection, "myUserVar")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_uservars.html>


## PluginVars

**Signature**

```
PluginVars([string:plugin_name]): light_userdata:plugin_preferences
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

**Used by APC plugin for:** persisted plugin settings (IP/port of the Go bridge, button mode flags).

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## AddonVars

**Signature**

```
AddonVars(string:addon_name): light_userdata:addon_variables
```

**Help page title:** `AddonVars(string)`

**Description**

The **AddonVars** function returns a handle to the set of variables connected to a specific addon.

                **Restriction:**

                The addon variable set is not helpful at the moment.

**Arguments**

- **String**:

            The string needs to be the name of the addon.

**Return**

- **Handle**:

            The function returns a handle of the set of variables.

**Example**

This example prints information connected to the "Demo" addon variable set. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Stores the handle to a variable set connected to the addon named 'Demo'.
    local variableSet = AddonVars("Demo")
    -- Check if the return is nil and print an error message
    if variableSet == nil then
        ErrPrintf("The variable set does not exists")
        return
    end
    Printf("=============== START OF DUMP ===============")
    variableSet:Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_addonvars.html>


## SetVar

**Signature**

```
SetVar(light_userdata:variables, string:varname, value): boolean:success
```

**Help page title:** `SetVar(handle, string, value)`

**Used by APC plugin for:** writes plugin settings.

**Description**

The SetVar Lua function sets a value to a specific variable in a set of variables. To learn more about the variables in plugins, look at the Variable Functions topic.

          If the variable exists, then the value is overwritten. If it does not exist, then it is created with value.

**Arguments**

- **Handle**:

            The handle of variable set.

- **String**:

            The name of the variable. It needs to be in quotation marks.

- **Value**:

            The value can be a string, integer, or double.

**Return**

- **Boolean**:

  -
                True / 1: The variable was set.

  -
                False / 0: The variable was not set.

**Example**

This example sets a value to the variable called "myUserVar" in the set of user variables if it exists.

```lua
return function()
    -- Sets the value of user variable "myUserVar" to "Hello World" and store the returned boolean in a Lua variable
    local success = SetVar(UserVars(), "myUserVar", "Hello World")
    -- Prints the ressult
    if success then
        Printf("Variable is stored.")
    else
        Printf("Variable is NOT stored!")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_setvar.html>


## GetVar

**Signature**

```
GetVar(light_userdata:variables, string:varname): value
```

**Help page title:** `GetVar(handle, string)`

**Used by APC plugin for:** reads plugin settings.

**Description**

The GetVar Lua function returns the value of a specific variable in a set of variables. To learn more about the variables in plugins, look at the Variable Functions topic.

**Arguments**

- **Handle**:

 The handle of variable set.

- **String**:

 The name of the variable. It needs to be in quotation marks.

**Return**

- **Value**:

 This is the value of the variable.

If the variable does not exist, then nil is returned.

**Example**

This example returns the value of a variable called "myUserVar" in the set of user variables if it exists:

```lua
return function()
    -- Get the value from a user variable called "myUserVar" - assuming it already exists
    local varValue = GetVar(UserVars(), "myUserVar")
    -- Print en error feedback or the value of the variable
    if varValue == nil then
        Printf("Variable returns nothing!")
    else
        Printf("Variable value is: " .. varValue)
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getvar.html>


## GetVarVersion

**Signature**

```
GetVarVersion(light_userdata:variables, string:varname): integer:version
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## DelVar

**Signature**

```
DelVar(light_userdata:variables, string:varname): boolean:success
```

**Help page title:** `DelVar(handle, string)`

**Description**

The **DelVar** Lua function deletes a specific variable in a set of variables. To learn more about the variables in plugins, have a look at the Variable Functions topic.

**Arguments**

- **Handle**:

            The handle of variable set.

- **String**:

            The name of the variable. It needs to be in quotation marks.

**Return**

- **Boolean**:

  -
                True / 1: The variable was deleted.

  -
                False / 0: The variable was not deleted.

          If the variable does not exist, then false is also returned.

**Example**

This example deletes a variable called "myUserVar" in the set of user variables.

```lua
return function()
    -- Deletes the variable called 'myUserVar' in the 'UserVars' variable set.
    local success = DelVar(UserVars(), "myUserVar")
    -- Prints the outcome of the deletion outcome.
    if success then
        Printf("Variable is deleted.")
    else
        Printf("Variable is NOT deleted!")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_delvar.html>

---

[Back to index](README.md)
