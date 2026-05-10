# Command execution & feedback

Run command-line strings, get feedback, log to System Monitor, and call commands on remote stations.

## Functions

- [`Cmd`](#cmd) [APC]
- [`CmdIndirect`](#cmdindirect) [APC]
- [`CmdIndirectWait`](#cmdindirectwait) [APC]
- [`Echo`](#echo) [APC]
- [`ErrEcho`](#errecho)
- [`Printf`](#printf) [APC]
- [`ErrPrintf`](#errprintf) [APC]
- [`RemoteCommand`](#remotecommand) [APC]
- [`CallRealtimeLockedProtected`](#callrealtimelockedprotected)

---


## Cmd

**Signature**

```
Cmd(string:formatted_command[ ,light_userdata:undo], ...): string:command_execution_result ('Ok', 'Syntax Error', 'Illegal Command', ...)
```

**Help page title:** `Cmd(string[, handle])`

**Used by APC plugin for:** fires command-line strings (e.g. `Go+`, `Off`, `Page`) from the OSC bridge.

**Description**

The **Cmd** Lua function executes a command in the grandMA3 command line. It is executed in a Lua task - not the Main task (standard typed commands are run in the Main task). It is executed synchronously, and it blocks the Lua task while executing. This means that a bad command has the potential to block the system.

          Alternative functions are CmdIndirect() and CmdIndirectWait().

**Arguments**

- **String**:

            A string with the command to be executed in the command line. Do not add a please or enter to execute the command.

- **Handle** (optional):

            A handle to an undo (oops) list. Learn more in the CreateUndo topic.

- **...** (optional):

            Additional arguments relevant for the command.

**Return**

- **String**:

            A string is returned with the execution feedback known from the command line feedback

  - **OK**:

                Command executed.

  - **Syntax Error**:

                The command was not executed because of a syntax error.

  - **Illegal Command**:

                Command not executed because of some illegal command or action.

          The returned string does not need to be used.

**Example**

This example executes the command "ClearAll" in the command line.

```lua
return function()
    --Execute the command directly
    Cmd("ClearAll")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_cmd.html>


## CmdIndirect

**Signature**

```
CmdIndirect(string:command[, light_userdata:undo[, light_userdata:target]]): nothing
```

**Help page title:** `CmdIndirect(string[, handle[, handle]])`

**Used by APC plugin for:** queues commands without blocking the Lua task; useful when reacting to MIDI without holding the realtime lock.

**Description**

The **CmdIndirect** Lua function executes a command within the grandMA3 command line. It is executed asynchronously in the Main task. It does not block the Lua execution since it is not executed in the Lua Task.

**Arguments**

- **String**:

            A string with the command to be executed in the command line. Do not add a please or enter to execute the command.

- **Handle **(optional):

            A handle to an undo (oops) list. Learn more in the CreateUndo topic.

- **Handle **(optional):

            This is a handle for the target for the command. The target can be a specific screen. See the example below.

**Return**

This function does not return anything.

**Example**

This example prints "1" and "2" in the Command Line History and let the main task open the Configure Display pop-up on screen 2.

```lua
return function()
    --Print something
    Printf("1")
    --Use the 'CmdIndirect' to open a pop-up
    CmdIndirect("Menu DisplayConfig", nil, GetDisplayByIndex(2))
    --Print something else
    Printf("2")
end
```

          The Command Line History shows:

                OK :

                Call Plugin 49

                1

                2

                OK :

                Menu "DisplayConfig"

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_cmdindirect.html>


## CmdIndirectWait

**Signature**

```
CmdIndirectWait(string:command[, light_userdata:undo[, light_userdata:target]]): nothing
```

**Help page title:** `CmdIndirectWait(string[, handle[, handle]])`

**Used by APC plugin for:** queue-and-wait variant when the plugin needs sequential execution but no return value.

**Description**

The **CmdIndirectWait** Lua function executes a command within the grandMA3 command line. It does not block the Lua execution and is executed synchronously in the main task. Synchronous commands wait for the command to be executed before executing any following command.

**Arguments**

- **String**:

            A string with the command to be executed in the command line. Do not add a please or enter to execute the command.

- **Handle **(optional):

            A handle to an undo (oops) list. Learn more in the CreateUndo topic.

- **Handle **(optional):

            This is a handle for the target for the command. The target can be a specific screen. See the example below.

**Return**

This function does not return anything.

**Example**

This example prints "1" and "2" in the Command Line History and lets the main task open the Configure Display pop-up on screen 2.

```lua
return function()
    --Print something
    Printf("1")
    --Use the 'CmdIndirectWait' to open a pop-up
    CmdIndirectWait("Menu DisplayConfig", nil, GetDisplayByIndex(2))
    --Print something else
    Printf("2")
end
```

          The Command Line History shows:

                OK :

                Call Plugin 50

                1

                OK :

                Menu "DisplayConfig"

                2

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_cmdindirectwait.html>


## Echo

**Signature**

```
Echo(string:formatted_command ...): nothing
```

**Help page title:** `Echo(string)`

**Used by APC plugin for:** developer logging to the system monitor while debugging the plugin.

**Description**

The **Echo** Lua function prints a string in the System Monitor.

**Arguments**

- **String**:

 The string text to be printed to the System Monitor.

**Return**

This function does not return anything.

**Example**

This example prints "Hello World!" on the System Monitor:

```lua
return function()
    -- Prints 'Hellow World!' in the system monitor in yellow text.
    Echo("Hello World!")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_echo.html>


## ErrEcho

**Signature**

```
ErrEcho(string:formatted_command ...): nothing
```

**Help page title:** `ErrEcho(string)`

**Description**

The **ErrEcho** Lua function prints a red error message on the System Monitor.

**Arguments**

- **String**:

 The string text is to be printed to the System Monitor.

**Return**

This function does not return anything.

**Example**

This prints "This is a red error message!" on the system monitor:

```lua
return function()
    -- Prints an error message in the system monitor in red text.
    ErrEcho("This is an error message!")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_errecho.html>


## Printf

**Signature**

```
Printf(string:formatted_command ...): nothing
```

**Help page title:** `Printf(string)`

**Used by APC plugin for:** developer logging via Lua-style formatting.

**Description**

The Printf Lua function prints a string in the Command Line History and System Monitor.

Command Line History

The **Command Line History** window shows feedback from the system based in the user input.

Learn more in the Command Line History topic.

System Monitor

The **System Monitor** window shows what is happening at the station. This includes feedback on user commands. It is a log of the different things happening in the background. It also shows warnings, errors, and changes to the system.

Learn more in the System Monitor topic.

**Arguments**

- **String**:

 The string text to be printed to the Command Line History.

**Return**

This function does not return anything.

**Example**

This example prints "Hello World!" in the Command Line History:

```lua
return function()
    Printf("Hello World!")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_printf.html>


## ErrPrintf

**Signature**

```
ErrPrintf(string:formatted_command ...): nothing
```

**Help page title:** `ErrPrintf(string)`

**Used by APC plugin for:** error logging when MIDI/OSC parsing fails.

**Description**

The **ErrPrintf** Lua function prints a red error message in the Command Line History and System Monitor.

**Arguments**

- **String**:

 The string text to be printed to the Command Line History.

**Return**

This function does not return anything.

**Example**

This example prints "This is a red error message!" in the Command Line History and System Monitor:

```lua
return function()
    -- Prints an error message in the command line feedback in red text.
    ErrPrintf("This is an error message!")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_errprintf.html>


## RemoteCommand

**Signature**

```
RemoteCommand(string:ip, string:command): boolean:success
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

**Used by APC plugin for:** fires commands on a different MA3 station (e.g. tracking-backup).

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## CallRealtimeLockedProtected

**Signature**

```
CallRealtimeLockedProtected(function:name): result of function
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)

---

[Back to index](README.md)
