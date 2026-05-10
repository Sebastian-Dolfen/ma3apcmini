# Hooks, callbacks & timers

Subscribe to object-change events, wait on changes, manage hooks, and schedule recurring Lua timers.

## Functions

- [`HookObjectChange`](#hookobjectchange) [APC]
- [`PrepareWaitObjectChange`](#preparewaitobjectchange)
- [`Unhook`](#unhook) [APC]
- [`UnhookMultiple`](#unhookmultiple)
- [`DumpAllHooks`](#dumpallhooks)
- [`Timer`](#timer) [APC]

---


## HookObjectChange

**Signature**

```
HookObjectChange(function:callback, light_userdata:handle, light_userdata:plugin_handle[, light_userdata:target]): integer:hook_id
```

**Help page title:** `HookObjectChange(function, handle, handle[, handle])`

**Used by APC plugin for:** subscribes to cue/sequence changes so the APC LEDs update without polling.

**Description**

The **HookObjectChange **Lua function automatically calls a function when a grandMA3 object changes.

**Arguments**

- **Function**:

            This must be the name of a function. This function is triggered every time the provided grandMA3 object changes.

- **Handle**:

            This is the handle for the grandMA3 objects that should be monitored for changes. The triggered function passes this handle on as the first argument.

- **Handle**:

            The handle must be for the plugin creating this HookObjectChange - it is the handle for "this" plugin.

- **Handle **(optional):

            This optional handle is for an object that will be passed on to the triggered function (as the third argument).

**Return**

- **Integer**:

            The function returns an integer identifying the hook. This can be saved to unhook the object later.

                **Hint:**

                See also these related functions: DumpAllHooks, Unhook, UnhookMultiple.

**Example**

To call a function every time the content of the sequence pool changes, create a plugin with this code:

```lua
-- Get the handle to this Lua component.
local luaComponentHandle = select(4,...)

function Main()
    -- Get a handle to the sequence pool.
    local hookObject = DataPool().Sequences
    -- Get a handle to this plugin.
    local pluginHandle = luaComponentHandle:Parent()
    -- Create the hook and save the Hook ID.
    SequenceHookId = HookObjectChange(MySequencePoolCallback, hookObject, pluginHandle)
    -- Print the returned Hook ID.
    Printf("HookId: " .. SequenceHookId)
end

-- This function is called when there are changes in the sequence pool.
function MySequencePoolCallback(obj)
    Printf(tostring(obj.name) .. " changed!")
end

return Main
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_hookobjectchange.html>


## PrepareWaitObjectChange

**Signature**

```
PrepareWaitObjectChange(light_userdata:handle[ ,integer:change_level_threshold]): boolean:true or nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## Unhook

**Signature**

```
Unhook(integer:hook_id): nothing
```

**Help page title:** `Unhook(integer)`

**Used by APC plugin for:** tears down the change hook when the plugin reloads.

**Description**

The **Unhook **Lua function removes a hook.

          Hooks are an automatically triggered function that activates when a grandMA3 object changes. A hook can be created using the HookObjectChange function.

**Arguments**

- **Integer**:

            This must be the integer matching the hook that should be unhooked.

                **Hint:**

                All hooks can be listed using the DumpAllHooks function, but this does not reveal the corresponding hook integer ID. Use the UnhookMultiple function if the integer is unknown.

**Return**

This function does not return anything.

                **Hint:**

                See also these related functions: DumpAllHooks, HookObjectChange, UnhookMultiple.

**Example**

This example unhooks the hook created using the example in the HookObjectChange - please run that example before this one.

```lua
return function()
    -- Unhooks the specific Hook integer ID.
    Unhook(SequenceHookId)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_unhook.html>


## UnhookMultiple

**Signature**

```
UnhookMultiple(function:callback(can be nil), light_userdata:handle to target(can be nil), light_userdata:handle to context (can be nil)): integer:amount of removed hooks
```

**Help page title:** `UnhookMultiple(function, handle, handle)`

**Description**

The **UnhookMultiple **Lua function unhooks multiple hooks based on an input. This input acts like a filter to identify all the hooks that should be unhooked.

          The DumpAllHooks function can be used to list all the existing hooks in the system.

**Arguments**

- **Function** or **nil**:

            This must be the name of a triggered function or nil.

- **Handle **or **nil**:

            This must be the handle for the target object or nil.

- **Handle** or **nil**:

            The must be the handle for the context object or nil.

          The target and context names can be seen using the DumpAllHooks function.

**Return**

- **Integer**:

            The function returns an integer indicating how many hooks were unhooked.

                **Hint:**

                See also these related functions: DumpAllHooks, HookObjectChange, Unhook.

**Example**

This example unhooks all hooked related to the function created in the example for the HookObjectChange - please run the example from that topic before running this one.

```lua
return function ()
    -- Unhooks all hooks related to the "MySequencePoolCallback" function.
    local amount = UnhookMultiple(MySequencePoolCallback)
    -- Print how many hooks that were unhooked.
    Printf(amount .. " hook(s) were unhooked.")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_unhookmultiple.html>


## DumpAllHooks

**Signature**

```
DumpAllHooks(nothing): nothing
```

**Help page title:** `DumpAllHooks()`

**Description**

The **DumpAllHooks** function prints a list of the hooks in the system. The list is only shown in the System Monitor.

**Arguments**

This function does not accept any arguments.

**Return**

The function does not return anything. It does print a list in the system monitor.

        **Hint:**

          See also these related functions: HookObjectChange, Unhook, UnhookMultiple.

**Example**

This example prints the list of hooks in the system monitor.

            The system monitor shows what is happening at the station. This includes feedback on user commands. It logs the different things happening in the background. It also shows warnings, errors, and changes to the system. Learn more in the System Monitor topic.

```lua
return function()
    -- Dumps a list of all the hooks in the System Monitor.
    Printf("=============== START OF HOOK DUMP ===============")
    DumpAllHooks()
    Printf("================ END OF HOOK DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_dumpallhooks.html>


## Timer

**Signature**

```
Timer(function:timer_function, integer:delay_time, integer:max_count[, function:cleanup][, light_userdata:context object]): nothing
```

**Help page title:** `Timer(string, integer, integer[, string[, handle]])`

**Used by APC plugin for:** drives the periodic OSC heartbeat / LED refresh cycle.

**Description**

The **Timer** Lua function call a different function using a timer. The other function can be called multiple times using the timer interval.

**Arguments**

- **Function**:

          This is the name of the function that is called multiple times using the timer.

- **Integer**:

          This is the wait time between the calls. The value is in seconds.

- **Integer**:

          This is the number of times the function is called.

- **Function** | **nil** (optional):

          This is an optional argument that is the name of a function that is called when the Timer function is finished.

- **Handle** (optional):

          This is an optional argument for a handle to an object that is passed to the called function.

**Return**

This function does not return anything.

**Example**

This example prints a greeting three times and then calls a clean up function:

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_timer.html>

---

[Back to index](README.md)
