# Undo / desk / show state

Open / close undo lists and inspect desk lock, remote-call activity and the unsaved-show flag.

## Functions

- [`CreateUndo`](#createundo)
- [`CloseUndo`](#closeundo)
- [`DeskLocked`](#desklocked)
- [`RemoteCallRunning`](#remotecallrunning)
- [`NeedShowSave`](#needshowsave)

---


## CreateUndo

**Signature**

```
CreateUndo(string:undo_text): light_userdata:undo_handle
```

**Help page title:** `CreateUndo(string)`

**Description**

The **CreateUndo** Lua function returns a handle to a list of commands and function calls grouped in the same oops action.

Functions can be executed with a reference to the undo handle. This adds the function to the undo list.

Undo lists need to be closed using the CloseUndo function.

**Arguments**

- **String**:

 A text string must be added. It can be used to identify the undo list.

**Return**

- **Handle**:

 The function returns the handle to the undo list.

**Example**

This example creates an undo list, performs a series of commands being added to the undo list, and closes the undo list. Now, the series of commands can be oopsed with one oops command.

```lua
return function()
    -- Create the undo group.
    local MyNewUndo = CreateUndo("MySelection")
    -- Make some command line actions linked to the undo.
    Cmd("ClearAll", MyNewUndo)
    Cmd("Fixture 1", MyNewUndo)
    Cmd("Fixture 2", MyNewUndo)
    Cmd("Fixture 5", MyNewUndo)
    Cmd("Fixture 7", MyNewUndo)
    -- Closing the undo group and store it's return in a variable.
    local closeSuccess = CloseUndo(MyNewUndo)
    -- Print the feedback from the closing action - 1 = Success / 0 = Failure.
    if closeSuccess == false then
        ErrPrintf("The CloseUndo was not successful")
    elseif closeSuccess == true then
        Printf("The CloseUndo was successful")
    else
        Printf("The CloseUndo did not return a meaningful result")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_createundo.html>


## CloseUndo

**Signature**

```
CloseUndo(light_userdata:undo_handle): boolean:closed (true if was closed, false - if it's still in use)
```

**Help page title:** `CloseUndo(handle)`

**Description**

The **CloseUndo** Lua function closes an open undo list. The function returns a boolean indicating if the function succeeds.

          Undo lists need to be created to be closed. See more about this in the CreateUndo function.

**Arguments**

- **Handle**:

            The handle of a created undo list.

**Return**

- **Boolean**:

  -
                True: The undo list was closed.

  -
                False: The undo list is still in use and cannot be closed.

**Example**

This example creates an undo list, performs a series of commands that are added to the undo list, and closes the undo list. Now the series of commands can be oopsed with one oops command.

```lua
return function()
    --Create the undo object 
    local MyNewUndo = CreateUndo("MySelection")
    --Create command actions connected to the undo object
    Cmd("ClearAll", MyNewUndo)
    Cmd("Fixture 1", MyNewUndo)
    Cmd("Fixture 2", MyNewUndo)
    Cmd("Fixture 5", MyNewUndo)
    Cmd("Fixture 7", MyNewUndo)
    --Close the undo group and store it's return in a variable
    local closeSuccess = CloseUndo(MyNewUndo)
    --Print the feedback from the closing action - 1 = Success / 0 = Failure.
    if closeSuccess == false then
        ErrPrintf("The CloseUndo was not successful")
    elseif closeSuccess == true then
        Printf("The CloseUndo was successful")
    else
        Printf("The CloseUndo did not return a meaningful result")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_closeundo.html>


## DeskLocked

**Signature**

```
DeskLocked(nothing): boolean:desk_is_locked
```

**Help page title:** `DeskLocked()`

**Description**

The **DeskLocked** Lua function returns a boolean indicating if the station is locked.

**Arguments**

This function does not accept any arguments.

**Return**

- **Boolean**:

            The boolean indicates if the station is desk locked or not.

  - **True** (or 1): The station is locked.

  - **False** (or 0): The station is not locked.

**Example**

This example prints the boolean number indicating the "DeskLocked" status to the Command Line History.

```lua
return function()
    -- The DeskLocked() return is printed.
    Printf("The desk is locked: " .. tostring(DeskLocked()))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_desklocked.html>


## RemoteCallRunning

**Signature**

```
RemoteCallRunning(nothing): boolean:remotecall_is_running
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## NeedShowSave

**Signature**

```
NeedShowSave(nothing): boolean:need_show_save
```

**Help page title:** `NeedShowSave()`

**Description**

The **NeedShowSave** Lua function returns a boolean indicating if there are unsaved changes to the showfile.

**Arguments**

This function does not accept any arguments.

**Return**

- **Boolean**:

 The boolean returns True if there are unsaved changes to the show file. False indicates that the show file has not changed since the last save. These indications do not include changes to the playback state of the show.

**Example**

This example prints feedback indicating if the show file should be saved or not.

```lua
return function ()
    -- Check if the show should be saved.
    if NeedShowSave() then
        Printf("You should save your showfile.")
    else
        Printf("You do not need to save your showfile.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_needshowsave.html>

---

[Back to index](README.md)
