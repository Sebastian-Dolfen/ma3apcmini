# Object access (handles, addresses, root pools)

Resolve top-level pools, the current selection / user / executor, and convert between handles, addresses and integers.

## Functions

- [`CmdObj`](#cmdobj)
- [`Root`](#root) [APC]
- [`Pult`](#pult)
- [`DefaultDisplayPositions`](#defaultdisplaypositions)
- [`Patch`](#patch)
- [`FixtureType`](#fixturetype)
- [`ShowData`](#showdata)
- [`ShowSettings`](#showsettings)
- [`DataPool`](#datapool) [APC]
- [`MasterPool`](#masterpool) [APC]
- [`DeviceConfiguration`](#deviceconfiguration)
- [`Programmer`](#programmer)
- [`ProgrammerPart`](#programmerpart)
- [`Selection`](#selection)
- [`CurrentUser`](#currentuser)
- [`CurrentProfile`](#currentprofile)
- [`CurrentEnvironment`](#currentenvironment)
- [`CurrentScreenConfig`](#currentscreenconfig)
- [`CurrentExecPage`](#currentexecpage)
- [`SelectedSequence`](#selectedsequence) [APC]
- [`GetCurrentCue`](#getcurrentcue) [APC]
- [`SelectedTimecode`](#selectedtimecode)
- [`SelectedLayout`](#selectedlayout)
- [`SelectedTimer`](#selectedtimer)
- [`GetSelectedAttribute`](#getselectedattribute)
- [`SelectedFeature`](#selectedfeature)
- [`SelectedDrive`](#selecteddrive)
- [`GetExecutor`](#getexecutor) [APC]
- [`LoadExecConfig`](#loadexecconfig)
- [`SaveExecConfig`](#saveexecconfig)
- [`GetObject`](#getobject) [APC]
- [`ObjectList`](#objectlist) [APC]
- [`FromAddr`](#fromaddr) [APC]
- [`ToAddr`](#toaddr) [APC]
- [`IntToHandle`](#inttohandle) [APC]
- [`HandleToInt`](#handletoint) [APC]
- [`StrToHandle`](#strtohandle)
- [`HandleToStr`](#handletostr)
- [`IsObjectValid`](#isobjectvalid) [APC]
- [`ClassExists`](#classexists)
- [`IsClassDerivedFrom`](#isclassderivedfrom)
- [`GetClassDerivationLevel`](#getclassderivationlevel)
- [`RefreshLibrary`](#refreshlibrary)

---


## CmdObj

**Signature**

```
CmdObj(nothing): light_userdata:handle
```

**Help page title:** `CmdObj()`

**Description**

The **CmdObj** Lua function returns information about the command line object.

**Arguments**

This function does not have any arguments.

**Return**

- **Handle**:

            The function returns a handle to the command line object.

**Example**

This example uses the Dump() function on the command object. It lists all the properties and lists the children and some extra examples of how the command line object can be used:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    --Store the handle to the command object
    local cmd = CmdObj()
    --Print all information about the command object
    Printf("=============== START OF DUMP ===============")
    cmd:Dump()
    Printf("================ END OF DUMP ================")
    --Print some selected elements from the command object - this is currently not in the online manual
    Printf("Current text in the command line: " ..cmd.cmdtext)
    Printf("Current cmd edit object: " ..tostring(cmd.editobject and cmd.editobject:ToAddr()))
    Printf("Current cmd destination: " ..tostring(cmd.destination and cmd.destination:ToAddr()))
    Printf("Current user of the command line: " ..tostring(cmd.user and cmd.user:ToAddr()))
    Printf("Current profile of the command line: " ..tostring(cmd.profile and cmd.profile:ToAddr()))
    Printf("Current DMX readout: " ..cmd.dmxreadout)
    Printf("Current amount steps: " ..cmd.maxstep)
    Printf("Current selected object: " ..tostring(cmd:GetSelectedObject() and cmd:GetSelectedObject():ToAddr()))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_cmdobj.html>


## Root

**Signature**

```
Root(nothing): light_userdata:handle
```

**Help page title:** `Root()`

**Used by APC plugin for:** top-level handle when resolving objects by address.

**Description**

The **Root** Lua function returns a handle to the object at the root position.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

            The function returns a handle to the Root object.

          Example

          This simple example prints the information of the Root object in the Command Line History using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- The following prints the dump for the root object
    Printf("=============== START OF DUMP ===============")
    Root():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_root.html>


## Pult

**Signature**

```
Pult(nothing): light_userdata:handle
```

**Help page title:** `Pult()`

**Description**

The **Pult** Lua function returns a handle to the current "Pult" object at position Root/GraphicsRoot/PultCollect. The "Pult" object contains display and device information.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the pult object.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- The following prints the dump for the pult object
    Printf("=============== START OF DUMP ===============")
    Pult():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_pult.html>


## DefaultDisplayPositions

**Signature**

```
DefaultDisplayPositions(nothing): light_userdata:handle
```

**Help page title:** `DefaultDisplayPositions()`

**Description**

The **DefaultDisplayPositions** Lua function returns the handle of the conventional default display positions, which contains the first seven screens as children.

For example, whether the command line, view bar, and encoder/playback bar are displayed.

**Arguments**

This function does not have any arguments.

**Return**

- **Handle**:

 The function returns a handle to the command line object.

**Example**

This example prints all the information about display 1 (child 1 of the default displays) using the Dump()** **function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Store a handle to display 1 (child 1 of the default displays).
    local display1 = DefaultDisplayPositions():Children()[1]
    -- Dumps information about the display.
    Printf("=============== START OF DUMP ===============")
    display1:Dump()
    Printf("================ END OF DUMP ================")
end
```

 This example toggles the Control Bar for display 1 with the help of the DefaultDisplayPositions object:

```lua
return function()
    -- Store a handle to display 1 (child 1 of the default displays).
    local display1 = DefaultDisplayPositions():Children()[1]
    -- Toggles the 'ShowMainMenu' setting.
    display1.ShowMainMenu = not display1.ShowMainMenu
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_defaultdisplaypositions.html>


## Patch

**Signature**

```
Patch(nothing): light_userdata:handle
```

**Help page title:** `Patch()`

**Description**

The **Patch** Lua function returns a handle to the patch object.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the patch.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- This example dumps all information about the patch object
    Printf("=============== START OF DUMP ===============")
    Patch():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_patch.html>


## FixtureType

**Signature**

```
FixtureType(nothing): light_userdata:handle
```

**Help page title:** `FixtureType()`

**Description**

The FixtureType Lua function returns a handle to the fixture type. The function does not accept any arguments, but the function must be executed when the command line destination is at a fixture type. If the command line destination is not a valid fixture type, then the function returns "nil".

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle** or **nil**:

            The handle for the fixture type or nil.

**Example**

This example prints the information about the second fixture type in the show:

```lua
return function ()
    -- The function returns the handle to the fixture at the current command line destination.
    -- Change to the "FixtureType" destination.
    Cmd("ChangeDestination FixtureType")
    -- Change to the second fixture type in the show.
    Cmd("ChangeDestination 2")
    -- Dump information about the Fixture Type handle. 
    Printf("=============== START OF DUMP ===============")
    FixtureType():Dump()
    Printf("================ END OF DUMP ================")
    -- Return the command line destination to the Root.
    Cmd("ChangeDestination Root")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_fixturetype.html>


## ShowData

**Signature**

```
ShowData(nothing): light_userdata:handle
```

**Help page title:** `ShowData()`

**Description**

ShowData is an object-free function that returns a handle to the object at position Root/ShowData.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

            The function returns a handle to the ShowData object.

          Example

          This simple example prints the information of the ShowData object in the Command Line History using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- The following prints the dump for the object for the show data
    Printf("=============== START OF DUMP ===============")
    ShowData():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_showdata.html>


## ShowSettings

**Signature**

```
ShowSettings(nothing): light_userdata:handle
```

**Help page title:** `ShowSettings()`

**Description**

ShowSettings is an object-free function that returns a handle to the object at Root/ShowData/ShowSettings.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

            The function returns a handle to the ShowSettings object.

**Example**

This simple example prints the information of the ShowSettings object using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- The following prints the dump for the object for the show settings
    Printf("=============== START OF DUMP ===============")
    ShowSettings():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_showsettings.html>


## DataPool

**Signature**

```
DataPool(nothing): light_userdata:handle
```

**Help page title:** `DataPool()`

**Used by APC plugin for:** root handle to the show's data pools - used to walk Sequences, Groups, Presets, etc.

**Description**

The **DataPool** Lua function references the currently selected DataPool and is used to read or edit properties within the data pool.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The function returns a handle to the DataPool object.

**Example**

This example uses the Dump() function on the data pool object. Dump lists all the properties and lists the children. Finally, the example also prints the name of the first sequence in the data pool.

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Dumps information about the datapool object.
    Printf("=============== START OF DUMP ===============")
    DataPool():Dump()
    Printf("================ END OF DUMP ================")
    -- Prints the name of the first sequence.
    Printf("Name of sequence 1: " .. DataPool().Sequences[1].Name)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_datapool.html>


## MasterPool

**Signature**

```
MasterPool(nothing): light_userdata:handle
```

**Help page title:** `MasterPool()`

**Used by APC plugin for:** root handle to the master pools (executors, pages) the APC mini drives.

**Description**

The **MasterPool** Lua function returns the handle to the masters.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The function returns the handle to the pool of masters.

**Example**

This example prints the information of the MasterPool object. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Print all informatin about the MasterPool object
    Printf("=============== START OF DUMP ===============")
    MasterPool():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_masterpool.html>


## DeviceConfiguration

**Signature**

```
DeviceConfiguration(nothing): light_userdata:handle
```

**Help page title:** `DeviceConfiguration()`

**Description**

The **DeviceConfiguration** Lua function returns a handle to the DeviceConfiguration object.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the DeviceConfiguration.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- This example dumps all information about the DeviceConfiguration object.
    Printf("=============== START OF DUMP ===============")
    DeviceConfiguration():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_deviceconfiguration.html>


## Programmer

**Signature**

```
Programmer(nothing): light_userdata:handle
```

**Help page title:** `Programmer()`

**Description**

The **Programmer** Lua function references the current programmer object.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

            The function returns a handle to the Programmer object.

**Example**

This example uses the Dump() function on the programmer object:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- Dumps information about the programmer object.
    Printf("=============== START OF DUMP ===============")
    Programmer():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_programmer.html>


## ProgrammerPart

**Signature**

```
ProgrammerPart(nothing): light_userdata:handle
```

**Help page title:** `ProgrammerPart()`

**Description**

The **ProgrammerPart** Lua function references the current programmer part object.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

            The function returns a handle to the programmer part object.

**Example**

This example uses the Dump() function on the programmer part object:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- Dumps information about the current programmer part object.
    Printf("=============== START OF DUMP ===============")
    ProgrammerPart():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_programmerpart.html>


## Selection

**Signature**

```
Selection(nothing): light_userdata:handle
```

**Help page title:** `Selection()`

**Description**

The Selection Lua function returns a handle to the object holding the current selection of fixtures.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The function returns a handle to the Selection object.

**Example**

This example prints the information of the Selection object in the Command Line History using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- The following prints the dump for the object for the selection
    Printf("=============== START OF DUMP ===============")
    Selection():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_selection.html>


## CurrentUser

**Signature**

```
CurrentUser(nothing): light_userdata:handle
```

**Help page title:** `CurrentUser()`

**Description**

The **CurrentUser** Lua function returns a handle to the current user.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the user.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Dumps information about the current user
    Printf("=============== START OF DUMP ===============")
    CurrentUser():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_currentuser.html>


## CurrentProfile

**Signature**

```
CurrentProfile(nothing): light_userdata:handle
```

**Help page title:** `CurrentProfile()`

**Description**

The **CurrentProfile** Lua function returns a handle to the current users' profile.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the user profile.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Dumps information about the current executor page
    Printf("=============== START OF DUMP ===============")
    CurrentProfile():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_currentprofile.html>


## CurrentEnvironment

**Signature**

```
CurrentEnvironment(nothing): light_userdata:handle
```

**Help page title:** `CurrentEnvironment()`

**Description**

The **CurrentEnvironment** Lua function returns a handle to the current users' selected environment.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the environment.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Dumps information about the current environment
    Printf("=============== START OF DUMP ===============")
    CurrentEnvironment():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_currentenvironment.html>


## CurrentScreenConfig

**Signature**

```
CurrentScreenConfig(nothing): light_userdata:handle
```

**Help page title:** `CurrentScreenConfig()`

**Description**

The **CurrentScreenConfig** Lua function returns a handle to the current users' screen configuration.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the screen configuration.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Dumps information about the current screen configuration
    Printf("=============== START OF DUMP ===============")
    CurrentScreenConfig():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_currentscreenconfig.html>


## CurrentExecPage

**Signature**

```
CurrentExecPage(nothing): light_userdata:handle
```

**Help page title:** `CurrentExecPage()`

**Description**

The **CurrentEnvironment** Lua function returns a handle to the current users' selected executor page.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the executor page.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Dumps information about the current executor page
    Printf("=============== START OF DUMP ===============")
    CurrentExecPage():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_currentexecpage.html>


## SelectedSequence

**Signature**

```
SelectedSequence(nothing): light_userdata:handle
```

**Help page title:** `SelectedSequence()`

**Used by APC plugin for:** current sequence handle - used to read cue list metadata for LED feedback.

**Description**

The **SelectedSequence **Lua function returns the handle of the selected sequence.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle for the selected sequence.

**Example**

This example prints all information about the selected sequence in the Command Line History using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- The following prints the dump for the selected sequence object
    Printf("=============== START OF DUMP ===============")
    SelectedSequence():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_selectedsequence.html>


## GetCurrentCue

**Signature**

```
GetCurrentCue(nothing): light_userdata:handle
```

**Help page title:** `GetCurrentCue()`

**Used by APC plugin for:** current cue handle - source of cue colour & label sent back to the APC.

**Description**

The **GetCurrentCue** Lua function returns a handle to the last activated cue in the selected sequence.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the cue.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Dumps information about the last activated cue in the selected sequence
    Printf("=============== START OF DUMP ===============")
    GetCurrentCue():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getcurrentcue.html>


## SelectedTimecode

**Signature**

```
SelectedTimecode(nothing): light_userdata:handle
```

**Help page title:** `SelectedTimecode()`

**Description**

The **SelectedTimecode **Lua function returns the handle of the selected timecode object.

The selected timecode object is the Timecode show currently selected in the Timecodes pool. Learn more in the Timecodes topics.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle for the selected timecode object.

**Example**

This example prints all information about the selected timecode show in the Command Line History using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- The following prints the dump for the selected timecode object
    local myTimecodeShow = SelectedTimecode()
    if myTimecodeShow ~= nil then
        Printf("=============== START OF DUMP ===============")
        myTimecodeShow:Dump()
        Printf("================ END OF DUMP ================")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_selectedtimecode.html>


## SelectedLayout

**Signature**

```
SelectedLayout(nothing): light_userdata:handle
```

**Help page title:** `SelectedLayout()`

**Description**

The **SelectedLayout **Lua function returns the handle of the selected layout.

Layouts are two-dimensional drafts where it is possible to arrange fixtures, macros, groups, and other pool objects. Learn more in the Layout topics.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle for the selected layout.

**Example**

This example prints all information about the selected layout in the Command Line History using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- The following prints the dump for the selected layout object
    Printf("=============== START OF DUMP ===============")
    SelectedLayout():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_selectedlayout.html>


## SelectedTimer

**Signature**

```
SelectedTimer(nothing): light_userdata:handle
```

**Help page title:** `SelectedTimer()`

**Description**

The **SelectedTimer **Lua function returns the handle of the selected timer object.

The selected timer object is the Timer currently selected in the Timers pool. Timers are stopwatch and timers that can be used to measure time. Learn more in the Timers topics.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle for the selected timer object.

**Example**

This example prints all information about the selected timer in the Command Line History using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- The following prints the dump for the selected timer object
    local myTimer = SelectedTimer()
    if myTimer ~= nil then
        Printf("=============== START OF DUMP ===============")
        myTimer:Dump()
        Printf("================ END OF DUMP ================")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_selectedtimer.html>


## GetSelectedAttribute

**Signature**

```
GetSelectedAttribute(nothing): light_userdata:handle
```

**Help page title:** `GetSelectedAttribute()`

**Description**

The **GetSelectedAttribute** Lua function returns a handle to the currently selected attribute.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle to the selected attribute.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- This example dumps all information about the currently selected attribute
    Printf("=============== START OF DUMP ===============")
    GetSelectedAttribute():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getselectedattribute.html>


## SelectedFeature

**Signature**

```
SelectedFeature(nothing): light_userdata:handle
```

**Help page title:** `SelectedFeature()`

**Description**

The **SelectedFeature **Lua function returns the handle of the selected feature.

**Arguments**

This function does not accept any arguments.

**Return**

- **Handle**:

 The returned handle for the selected feature.

**Example**

This example prints all information about the selected feature in the Command Line History using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- The following prints the dump for the selected feature object
    Printf("=============== START OF DUMP ===============")
    SelectedFeature():Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_selectedfeature.html>


## SelectedDrive

**Signature**

```
SelectedDrive(nothing): light_userdata:handle
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetExecutor

**Signature**

```
GetExecutor(integer:executor): light_userdata:executor, light_userdata:page
```

**Help page title:** `GetExecutor(integer)`

**Used by APC plugin for:** resolves the currently selected executor + page so the plugin can light its row of buttons.

**Description**

The **GetExecutor **Lua function returns the handles of the executor and the page based on the executor number.

**Arguments**

- **Integer**:

 The integer number for the executor.

**Return**

- **Handle - Executor**:

 The returned handle to the executor.

- **Handle - Page**:

 The returned handle to the page.

**Example**

This example stores the handles for executor number 201. It then uses the Dump() function to show the data for the two handles.

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- This saves the handles for executor 201 on the selected page.
    local executorHandle, pageHandle = GetExecutor(201)
    -- exit the function and print an error message if any of the handles are nil.
    if executorHandle == nil or pageHandle == nil then
        ErrPrintf("There is not a valid object on executor 201, please assign something and try again.")
        return
    end
    -- The following prints the dumps of the two handles.
    Printf("============ START OF EXEC DUMP =============")
    executorHandle:Dump()
    Printf("================ END OF DUMP ================")
    Printf("============ START OF PAGE DUMP =============")
    pageHandle:Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getexecutor.html>


## LoadExecConfig

**Signature**

```
LoadExecConfig(light_userdata:executor): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SaveExecConfig

**Signature**

```
SaveExecConfig(light_userdata:executor): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetObject

**Signature**

```
GetObject(string:address): light_userdata:handle
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

**Used by APC plugin for:** resolves an address (e.g. `Sequence 1`) to a handle for inspection.

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## ObjectList

**Signature**

```
ObjectList(string:address[, {['selected_as_default'=boolean:enabled], ['reverse_order'=boolean:enabled]}): {light_userdata:handles}
```

**Help page title:** `ObjectList(string[, table])`

**Used by APC plugin for:** enumerates pools (e.g. all sequences) when building an APC layout from the show.

**Description**

The **ObjectList** Lua function returns a table with handles. The table is created based on a string input that should create a selection.

**Arguments**

- **String**:

          The string must be a command that would create a range of objects in the command line.

- **Table** (optional):

          The table can contain two possible named elements. Each element can have a boolean true or false value. See the examples below for how to use them.

  - **'reverse_order'**:

              This must have a boolean value. If this is true then the returned list is in reverse order.

  - **'selected_as_default'**:

              This must have a boolean value. If this is true then the object list will only contain the object that is selected in the pool. For instance, it only returns the currently selected filter from the filter pool.

**Return**

- **Table**:

          The function returns a table with handles to the objects based on the string argument.

**Example**

This example returns the names and patch addresses of fixtures 1 through 10. It assumes these fixtures exist - if they do not, then it returns an error text.

```lua
return function()
    -- Create a list of handles based on the "Fixture 1 Thru 10" selection and store it in a table. 
    local myObjects = ObjectList("Fixture 1 Thru 10", {reverse_order=true})
    -- If the selection returned a table, then go through all elements and print information of the object.
    if myObjects~= nil then
        for i in pairs(myObjects) do
            Printf("Fixture: " .. myObjects[i].name .. " - Patch: " ..myObjects[i].patch)
        end
    else
        ErrPrintf("An error occured. Does Fixture 1 Thru 10 exist?")
    end
end
```

        This example creates an object list with the selected sequence. It then dumps all information about the sequence using the Dump function.

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Create a list of one handle to the selected sequence and store it to a table. 
    local myObjects = ObjectList("Sequence", {selected_as_default=true})
    -- If the selection returned a table, then dump the first (and only) element.
    if myObjects~= nil then
        myObjects[1]:Dump()
    else
        ErrPrintf("An error occured.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_objectlist.html>


## FromAddr

**Signature**

```
FromAddr(string:address[, light_userdata:base_handle]): light_userdata:handle
```

**Help page title:** `FromAddr(string[, handle])`

**Used by APC plugin for:** address-to-handle conversion when scripts pass string addresses across IPC.

**Description**

The **FromAddr** Lua function converts a numbered string address into a handle that can be used in commands.

**Arguments**

- **String**:

            A text string identifying an object. It can be a numbered or named address.

- **Handle** (optional):

            The default is to write the address from the root location. This optional handle can specify a different base location. It still needs to be a base location in the address path from the root to the object.

**Return**

- **Handle**:

            The handle for the addressed object.

**Example**

This example prints the address of the first sequence:

```lua
return function()
    -- Converts the string to a handle and store in a variabel.
    local mySequenceHandle = FromAddr("14.14.1.6.1")
    -- Converts the handle back to a numbered string and prints it.
    Printf("The address is: " ..mySequenceHandle:Addr())
    -- Converts the handle to a named string and prints it.
    Printf("The address is: " ..mySequenceHandle:AddrNative())
    
    -- Store the handle of the selected datapool.
    local myDataPool = DataPool()
    -- Prints the address of the selected datapool.
    Printf("The datapool address is: " ..myDataPool:Addr())
    
    --- The follwoing example uses the name of a sequence in the sequence pool. 
    --- Please adjust the "Default" name in the next line to match an existing named sequence.
    -- Finds the address based on the base location and a text string with names.
    local alsoMySequenceHandle = FromAddr("Sequences.Default", myDataPool)
    -- Converts the handle back to a numbered string and prints it.
    Printf("The address is: " ..alsoMySequenceHandle:Addr())
    -- Converts the handle to a named string and prints it.
    Printf("The address is: " ..alsoMySequenceHandle:AddrNative())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_fromaddr.html>


## ToAddr

**Signature**

```
ToAddr(light_userdata:handle, boolean:with_name[, boolean:use_visible_addr]): string:address
```

**Help page title:** `ToAddr(handle[, boolean])`

**Used by APC plugin for:** handle-to-address conversion for human-readable logging back over OSC.

**Description**

The **ToAddr** Lua object-free function converts a handle to an address string that can be used in commands.

          See the Handle topic for more info regarding handles, addresses, and links to other related functions.

**Arguments**

- **Handle**:

            The function takes a handle of an object as an argument.

- **Boolean** (optional):

            This returns the address using the names instead of numbers. The default is False, which returns the number version of the address.

**Return**

- **String**:

            String with the address value.

**Example**

This example prints the address of the selected sequence in both the numbered and named versions.

```lua
return function ()
    local mySequence = SelectedSequence()
    -- Print the address to the selected sequence in number and name format.
    Printf(ToAddr(mySequence))
    Printf(ToAddr(mySequence, true))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_toaddr.html>


## IntToHandle

**Signature**

```
IntToHandle(integer:handle): light_userdata:handle
```

**Help page title:** `IntToHandle(integer)`

**Used by APC plugin for:** rebuilds handles received as integers across the Lua/Go bridge.

**Description**

The **IntToHandle** Lua function converts an integer number into a handle. The integer needs to correlate with an actual handle.

See the Handle topic for more info regarding handles and links to other related functions.

**Arguments**

- **Integer**:

 The integer that correlates to an object's handle.

**Return**

- **Handle**:

 The returned handle of the object correlates with the integer.

**Example**

This example prints the handle integer number for the selected sequence. It also converts the integer back to a handle and uses this to print the name of the sequence:

```lua
return function()
    -- Convert the handle of the currently selected sequence to an integer
    local handleInt = HandleToInt(SelectedSequence())
    -- Print the handle integer
    Printf("The handle integer number of the selected sequence: %i", HandleToInt(SelectedSequence()))
    -- Convter the integer back to a hanndle and use it to get the sequence name
    Printf("The name of the selected sequence is: %s", IntToHandle(handleInt).name)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_inttohandle.html>


## HandleToInt

**Signature**

```
HandleToInt(light_userdata:handle): integer:handle
```

**Help page title:** `HandleToInt(handle)`

**Used by APC plugin for:** serialises a handle for transport via OSC/JSON.

**Description**

The **HandleToInt** Lua function converts a handle into an integer format.

See the Handle topic for more info regarding handles and links to other related functions.

**Arguments**

- **Handle**:

 The handle of the object.

**Return**

- **Integer**:

 The returned integer is the handle converted to an integer.

**Example**

This example prints the handle integer number for the selected sequence. It also converts the integer back to a handle and uses this to print the name of the sequence:

```lua
return function()
    Printf("The integer number for the handle of the selected sequence: %i", HandleToInt(SelectedSequence()))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_handletoint.html>


## StrToHandle

**Signature**

```
StrToHandle(string:handle(in H#... format)): light_userdata:handle
```

**Help page title:** `StrToHandle(string)`

**Description**

The object-free StrToHandle Lua function converts a string with a hexadecimal number format into a handle. The string needs to correlate with an actual handle.

          See the Handle topic for more info regarding handles and links to other related functions.

**Arguments**

- **String**:

            The string with a handle number in a hexadecimal format.

**Return**

- **Handle**:

            The returned handle based on the string with a hexadecimal number.

**Example**

This example prints the handle hex number for the selected sequence. It also converts the string back to a handle and uses this to print the name of the sequence:

```lua
return function()
    -- Store a variable with the string of the handle converted to hex
    local mySeqStr = HandleToStr(SelectedSequence())
    -- Print some feedback with the handle in a string version
    Printf("The handle for the selected sequence (string version): %s", mySeqStr)
    -- Print some feedback where the string is converted back to a handle
    Printf("The name of the selected sequence is: %s", StrToHandle(mySeqStr).name)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_strtohandle.html>


## HandleToStr

**Signature**

```
HandleToStr(light_userdata:handle): string:handle(in H#... format)
```

**Help page title:** `HandleToStr(handle)`

**Description**

The **HandleToStr** Lua function converts a handle into a string in a hexadecimal number format.

          See the Handle topic for more info regarding handles and links to other related functions.

**Arguments**

- **Handle**:

            The handle of the object.

**Return**

- **String**:

            The returned string is the handle number converted to a hexadecimal format.

**Example**

This example prints the handle hex number for the selected sequence. It also converts the string back to a handle and uses this to print the name of the sequence:

```lua
return function()
    Printf("The string (in hex format with 'H#' in front) for the handle of the selected sequence: %s",HandleToStr(SelectedSequence()))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_handletostring.html>


## IsObjectValid

**Signature**

```
IsObjectValid(light_userdata:handle): boolean:valid
```

**Help page title:** `IsObjectValid(handle)`

**Used by APC plugin for:** guards against stale handles after pool deletions.

**Description**

The **IsObjectValid **function returns a boolean true or nil depending on whether the supplied argument is a valid object.

**Arguments**

- **Handle**:

 The argument should be the handle to a possible object.

**Return**

- **Boolean or nil**:

 The returned value is a boolean True if the handle is a valid object or it returns nil if it is not a valid object.

**Example**

This example below examines if "Root()" is a valid object and prints meaningful feedback:

```lua
return function()
    --Create a variable with the possible object
    local myObject = Root()
    --Check if it is an object
    local myReturn = IsObjectValid(myObject)
    --Print the result
    if myReturn == nil then
        ErrPrintf("It is not a valid object")
    else
        Printf("It is an object")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_isobjectvalid.html>


## ClassExists

**Signature**

```
ClassExists(string:class_name): boolean:result
```

**Help page title:** `ClassExists(string)`

**Description**

The **ClassExists** Lua function returns a boolean indicating whether the provided string is a class.

**Arguments**

- **String**:

 A string containing a single word that could be a class.

**Return**

- **Boolean**:

 The function returns a boolean.

  - **True**:

 The provided word is a class.

  - **False**:

 The provided input is not a class.

**Example**

This example asks if the word "Display" is a class and returns proper feedback.

```lua
return function()
    -- Store a string with the class name
    local className = "Display"
    -- Check if the class exists and then provide proper feedback
    if ClassExists(className) then
        Printf("The class '%s' exists", className)
    else
        Printf("The class '%s' does not exists", className)
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_classexists.html>


## IsClassDerivedFrom

**Signature**

```
IsClassDerivedFrom(string:derived_name, string:base_name): boolean:result
```

**Help page title:** `IsClassDerivedFrom(string, string)`

**Description**

The **IsClassDerivedFrom** Lua function returns a boolean indicating if a class is derived from a different class.

**Arguments**

- **String**:

 This string needs to be the name of the class that might be derived from a different class.

- **String**:
This string needs to be the name of the class that might be the base class.

**Return**

- **Boolean**:

 The returned boolean indicates if the class is derived from the base class.

**Example**

This example checks if a class is derived from a different class and returns useful feedback.

```lua
return function()
    -- Set the value of the two strings.
    local derivedName = "World"
    local baseName = "Group"
    -- Check if the derivedName is the name of a class derived from the baseName class.
    local isDerived = IsClassDerivedFrom(derivedName, baseName)
    -- Provide feedback.
    if isDerived then
        Printf(derivedName .. " is derived from " .. baseName)
    else
        Printf(derivedName .. " is not derived from " .. baseName)
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_isclassderivedfrom.html>


## GetClassDerivationLevel

**Signature**

```
GetClassDerivationLevel(string:class_name): integer:result or nothing
```

**Help page title:** `GetClassDerivationLevel(string)`

**Description**

The **GetClassDerivationLevel** Lua function returns an integer indicating the derivation level index for a class based on a class name.

**Arguments**

- **String**:

 This string needs to be the name of a class.

**Return**

- **Integer**:

 The returned integer indicates the class derivation level.

**Example**

This example prints the index integer for the Pool class in the Command Line History:

```lua
return function()
    -- Get the index integer for the "Pool" class.
    local classDerivationLevel = GetClassDerivationLevel("Pool")
    -- Create a valid Printf return.
    if classDerivationLevel == nil then
        Printf("The return is nil")
    else
        Printf("The ClassDerivationLevel index for 'Pool' is: %i", classDerivationLevel)
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getclassderivationlevel.html>


## RefreshLibrary

**Signature**

```
RefreshLibrary(light_userdata:handle): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)

---

[Back to index](README.md)
