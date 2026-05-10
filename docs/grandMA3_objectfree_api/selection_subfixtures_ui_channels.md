# Selection, subfixtures & UI channels

Walk the current selection, enumerate subfixtures, and inspect UI / RT channels, attributes and channel functions.

## Functions

- [`SelectionTable`](#selectiontable)
- [`ChannelTable`](#channeltable)
- [`SelectionFirst`](#selectionfirst)
- [`SelectionNext`](#selectionnext)
- [`SelectionCount`](#selectioncount)
- [`SelectionComponentX`](#selectioncomponentx)
- [`SelectionComponentY`](#selectioncomponenty)
- [`SelectionComponentZ`](#selectioncomponentz)
- [`GetSubfixtureCount`](#getsubfixturecount)
- [`GetSubfixture`](#getsubfixture)
- [`GetUIChannelCount`](#getuichannelcount)
- [`GetRTChannelCount`](#getrtchannelcount)
- [`GetAttributeCount`](#getattributecount)
- [`GetUIChannels`](#getuichannels)
- [`GetRTChannels`](#getrtchannels)
- [`GetUIChannel`](#getuichannel)
- [`GetRTChannel`](#getrtchannel)
- [`GetAttributeByUIChannel`](#getattributebyuichannel)
- [`FirstDmxModeFixture`](#firstdmxmodefixture)
- [`NextDmxModeFixture`](#nextdmxmodefixture)
- [`GetAttributeIndex`](#getattributeindex)
- [`GetUIChannelIndex`](#getuichannelindex)
- [`GetChannelFunctionIndex`](#getchannelfunctionindex)
- [`GetChannelFunction`](#getchannelfunction)
- [`GetTokenName`](#gettokenname)
- [`GetTokenNameByIndex`](#gettokennamebyindex)
- [`SelectionNotifyBegin`](#selectionnotifybegin)
- [`SelectionNotifyObject`](#selectionnotifyobject)
- [`SelectionNotifyEnd`](#selectionnotifyend)

---


## SelectionTable

**Signature**

```
SelectionTable(nothing): table of subfixture_index
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## ChannelTable

**Signature**

```
ChannelTable(string:attribute_name or integer:attribute_index): table of ui_channel_index
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SelectionFirst

**Signature**

```
SelectionFirst(nothing): integer:first_subfixture_index, integer:x, integer:y, integer:z
```

**Help page title:** `SelectionFirst()`

**Description**

The **SelectionFirst** Lua function returns a set of integers for the selection's first fixture. It is the patch index number and the XYZ grid values in the selection grid.

It is not required to use all four returned integers, but they are returned in order.

**Arguments**

This function does not accept any arguments.

**Return**

- **Integer**:

 The returned number is the patch index of the first fixture in the current selection. It is not the FID or CID. The index is 0-based.

- **Integer**:

 The returned number is the current position on the X-axis in the selection grid. The selection grid is 0-based.

- **Integer**:

 The returned number is the current position on the Y-axis in the selection grid. The selection grid is 0-based.

- **Integer**:

 The returned number is the current position on the Z-axis in the selection grid. The selection grid is 0-based.

**Example**

This example prints the returned numbers of the first fixture in the selection, to the Command Line History:

```lua
return function()
    -- Store the return in a local variable
    local fixtureIndex, gridX, gridY, gridZ = SelectionFirst();

    -- Cancel the plugin if no fixture is selected
    assert(fixtureIndex,"Please select a fixture and try again.");

    -- Print the index number of the first fixture in the selection
    Printf("First selected fixture has index number: "..fixtureIndex
        .." and gridX value: "..gridX
        .." and gridY value: "..gridY
        .." and gridZ value: "..gridZ);
end
```

Related Functions

- SelectionNext

- GetSubfixture

- GetSubFixtureCount

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_selectionfirst.html>


## SelectionNext

**Signature**

```
SelectionNext(integer:current_subfixture_index): integer:next_subfixture_index, integer:x, integer:y, integer:z
```

**Help page title:** `SelectionNext()`

**Description**

The SelectionNext function returns a set of integers for the next fixture in a selection based on the index number input as an argument. It is the index number in the patch and the XYZ grid values in the selection grid.

It is not required to use all four returned integers, but they are returned in order.

**Arguments**

- **Integer**:

 The index number is used to find the next fixture. The index number needs to be part of the current selection.

**Return**

- **Integer**:

 The returned number is the patch index of the first fixture in the current selection. It is not the FID or CID. The index is 0-based.

- **Integer**:

 The returned number is the current position on the X-axis in the selection grid. The selection grid is 0-based.

- **Integer**:

 The returned number is the current position on the Y-axis in the selection grid. The selection grid is 0-based.

- **Integer**:

 The returned number is the current position on the Z-axis in the selection grid. The selection grid is 0-based

**Example**

This example prints the patch index number and grid positions of all the fixtures in the current selection:

```lua
return function()
    -- Store the return in a local variable
    local fixtureIndex, gridX, gridY, gridZ = SelectionFirst()

    -- Cancel the plugin if no fixture is selected
    assert(fixtureIndex,"Please select a (range of) fixture(s) and try again.")
    
    -- Loop that prints the index and gridpositions of all the fixtures in the selection 
    while fixtureIndex do
        Printf('The fixture has index number: %i and gridposition %i / %i / %i',
          fixtureIndex, gridX, gridY, gridZ);
        
        -- Here is SelectionNext actually used to find the next fixture in the selection
        fixtureIndex, gridX, gridY, gridZ = SelectionNext(fixtureIndex)
    end
end
```

Related Functions

- SelectionFirst

- GetSubfixture

- GetSubfixtureCount

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_selectionnext.html>


## SelectionCount

**Signature**

```
SelectionCount(nothing): integer:amount_of_selected_subfixtures
```

**Help page title:** `SelectionCount()`

**Description**

The **SelectionCount **Lua function returns a number indicating the total amount of currently selected fixtures.

**Arguments**

This function does not accept any arguments.

**Return**

- **Integer**:

 The function returns an integer number depicting the total amount of fixtures in the current selection.

 If there is no selection, then it returns 0.

**Example**

This example prints the number of fixtures in the current selection to the Command Line History:

```lua
return function()
    Printf('Number of fixtures in the current selection: %i', SelectionCount())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_selectioncount.html>


## SelectionComponentX

**Signature**

```
SelectionComponentX(nothing): integer:min, integer:max, integer:index, integer:block, integer:group
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SelectionComponentY

**Signature**

```
SelectionComponentY(nothing): integer:min, integer:max, integer:index, integer:block, integer:group
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SelectionComponentZ

**Signature**

```
SelectionComponentZ(nothing): integer:min, integer:max, integer:index, integer:block, integer:group
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetSubfixtureCount

**Signature**

```
GetSubfixtureCount(nothing): integer:subfixture_count
```

**Help page title:** `GetSubfixtureCount()`

**Description**

The **GetSubfixtureCount **Lua function returns the total number of fixtures that are patched within the show file.

**Arguments**

This function does not accept any arguments.

**Return**

- **Integer**:

 The returned integer number represents the total amount of patched fixtures on all the stages in the show file.

**Example**

This example prints the total number of patched fixtures in the Command Line History:

```lua
return function ()
    Printf('Total number of patched fixtures: %i', GetSubfixtureCount())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getsubfixturecount.html>


## GetSubfixture

**Signature**

```
GetSubfixture(integer:subfixture_index): light_userdata:subfixture
```

**Help page title:** `GetSubfixture(integer)`

**Description**

The **GetSubfixture** Lua function returns the handle of the fixture specified by its patch index number.

**Arguments**

- **Integer**:

 The patch index number for a fixture. This is also known as the "subfixtureindex".

**Return**

- **Handle**:

 The function returns a handle to the fixture object matching the provided index number.

**Example**

This example uses a fixture selection to print all the information (in the Command Line History) about the first fixture in the selection using the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- Check for a fixture selection, by returning an index for the first fixture
    if (SelectionFirst()) then
        -- There is a fixture selection, store the index for the first fixture 
        local fixtureIndex = SelectionFirst()
        -- Dump all information about the fixture
        Printf("=============== START OF DUMP ===============")
        GetSubfixture(fixtureIndex):Dump()
        Printf("================ END OF DUMP ================")
    else
        -- There needs to be a selection of at least one fixture
        Printf("Please select a fixture")
    end
end
```

Related Functions

- SelectionFirst

- SelectionNext

- GetSubfixtureCount

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getsubfixture.html>


## GetUIChannelCount

**Signature**

```
GetUIChannelCount(nothing): integer:ui_channel_count
```

**Help page title:** `GetUIChannelCount()`

**Description**

The **GetUIChannelCount **Lua function returns a number indicating the total amount of UI channels.

**Arguments**

This function does not accept any arguments.

**Return**

- **Integer**:

 The function returns an integer number depicting the total amount of UI channels.

**Example**

This example prints the number of UI channels to the Command Line History:

```lua
return function()
    Printf("The number of UI channels is " .. GetUIChannelCount())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getuichannelcount.html>


## GetRTChannelCount

**Signature**

```
GetRTChannelCount(nothing): integer:rt_channel_count
```

**Help page title:** `GetRTChannelCount()`

**Description**

The **GetRTChannelCount **Lua function returns a number indicating the total amount of RT channels.

**Arguments**

This function does not accept any arguments.

**Return**

- **Integer**:

 The function returns an integer number depicting the total amount of RT channels.

**Example**

This example prints the number of RT channels to the Command Line History:

```lua
return function()
    Printf("The number of RT channels is " .. GetRTChannelCount())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getrtchannelcount.html>


## GetAttributeCount

**Signature**

```
GetAttributeCount(nothing): integer:attribute_count
```

**Help page title:** `GetAttributeCount()`

**Description**

The **GetAttributeCount** Lua function returns the total number of attribute definitions in the show.

**Arguments**

This function does not accept any arguments.

**Return**

- **Integer**:

            The returned integer number represents the total amount of attribute definitions in the show file.

**Example**

This example prints the returned number in the Command Line History.

```lua
return function()
    Printf("Attribute count is %i", GetAttributeCount())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getattributecount.html>


## GetUIChannels

**Signature**

```
GetUIChannels(integer:subfixture_index or light_userdata: subfixture_handle[, boolean:return_as_handles]): {integer:ui_channels} or {light_userdata:ui_channels}
```

**Help page title:** `GetUIChannels(integer[,boolean] OR handle[,boolean])`

**Description**

The **GetUIChannels** Lua function returns a table with UI Channel indexes or a table with handles to the UI Channel objects. There are two different types of arguments for this function.

**Arguments**

- **Integer**:

 The integer should be the index number for a (sub)fixture.

- **Boolean **(Optional):

  - **True**:

 The returned table contains handles for UI Channel objects.

  - **False **(default):

 The returned table contains integer index values to the UI Channel objects.

- OR -

- **Handle**:

 The handle should relate to a (sub)fixture object.

- **Boolean** (Optional):

  - **True**:

 The returned table contains handles for UI Channel objects.

  - **False **(default):

 The returned table contains integer index values to the UI Channel objects.

**Return**

- **Table**:

 The returned table can be a list of UI Channel indexes or handles to the same UI Channel indexes.

**Example**

Example 1

This example prints a list of UI Channel indexes for the first fixture in the selection. It uses an index number as input:

```lua
return function()
    -- Creates a table of indexes of the UI channels of the first selected fixture.
    local uiChannels = GetUIChannels(SelectionFirst())
    if uiChannels == nil then
        ErrPrintf("Please select a fixture and try again")
        return
    end
    for key,value in ipairs(uiChannels) do
        Printf("List index number ".. key .. " :  UIChannel Index = " .. value)
    end
end
```

Example 2

This example prints a list of UI Channel indexes and attributes for the first fixture in the selection. It uses a handle as the input:

```lua
return function()
    local fixtureHandle = GetSubfixture(SelectionFirst())
    -- Creates a table of handles to the UI channels of the first selected fixture.
    local uiChannels = GetUIChannels(fixtureHandle, true)
    if uiChannels == nil then
        ErrPrintf("Please select a fixture and try again")
        return
    end
    for key,value in pairs(uiChannels) do
        Printf("List index number ".. key .. ": UIChannel Index = %i, (Sub)Attribute = %s", value.INDEX-1, value.SUBATTRIBUTE)
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getuichannels.html>


## GetRTChannels

**Signature**

```
GetRTChannels(integer:fixture index or light_userdata: reference_to_fixture_object[, boolean:return_as_handles]): {integer:rt_channels} or {light_userdata:rt_channels}
```

**Help page title:** `GetRTChannels(integer[,boolean] OR handle[,boolean])`

**Description**

The **GetRTChannels** Lua function returns a table with RT Channel indexes or a table with handles to the RT Channel objects. There are two different types of arguments for this function.

**Arguments**

- **Integer**:

 The integer should be the index number for a (sub)fixture.

- **Boolean **(Optional):

  - **True**:

 The returned table contains handles for RT Channel objects.

  - **False **(default):

 The returned table contains integers index values to the RT Channel objects.

- OR -

- **Handle**:

 The handle should relate to a (sub)fixture object.

- **Boolean** (Optional):

  - **True**:

 The returned table contains handles for RT Channel objects.

  - **False **(default):

 The returned table contains integers index values to the RT Channel objects.

**Return**

- **Table**:

 The returned table can be a list of RT Channel indexes or handles to the same RT Channels.

**Example**

Example 1

This example prints a list of RT Channel indexes for the first fixture in the selection. It uses an index number as input:

```lua
return function()
    -- Get the index number for the first fixture in the current selection 
    local fixtureIndex = SelectionFirst()
    -- Get the indexes of the RT channels
    local rtChannels = GetRTChannels(fixtureIndex, false)
    -- Print an error message if returnd table is nil
    if rtChannels == nil then
        ErrPrintf("Please select a fixture and try again")
        return
    end
    -- Print the table content
    for key,value in ipairs(rtChannels) do
        Printf("List index number ".. key .." : RTChannel index number = ".. value)
    end
end
```

Example 2

This example prints a list of RT Channel indexes and attributes for the first fixture in the selection. It uses a handle as the input:

```lua
return function()
    -- Get a handle to the first fixture in the current selection
    local fixtureHandle = GetSubfixture(SelectionFirst())
    if fixtureHandle == nil then
        ErrPrintf("Please select a fixture and try again")
        return
    end
    -- Creates a table of handles to the RT channels of the first selected fixture.
    local rtChannels = GetRTChannels(fixtureHandle, true)
    if rtChannels == nil then
        ErrPrintf("Please select a fixture and try again")
        return
    end
    -- Print DMX addresses of the RT Channels for the fixture
    for key,value in ipairs(rtChannels) do
        Printf("List index number ".. key .. ": RTChannel Index = %i, Coarse DMX addr. = %s, Fine DMX addr. = %s", value.INDEX, value.COARSE, value.FINE)
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getrtchannels.html>


## GetUIChannel

**Signature**

```
GetUIChannel(integer:ui_channel_index or light_userdata: subfixture_reference, integer:attribute_index or string:attribute_name): table:ui_channel_descriptor
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetRTChannel

**Signature**

```
GetRTChannel(integer:rt_channel_index): table:rt_channel_descriptor
```

**Help page title:** `GetRTChannel(integer)`

**Description**

The **GetRTChannel** Lua function returns a table with information about the related RT Channel.

**Arguments**

- **Integer**:

 The integer should be the index number for an RT Channel.

**Return**

- **Table**:

 The returned table contains related numbers, tables, and handles with a named identifying key:

  - handle "fixture"

  - handle "subfixture"

  - handle "dmx_channel

  - integer "dmx_default"

  - integer "dmx_highlight"

  - integer "dmx_lowlight"

  - integer "ui_index_first"

  - integer "rt_index"

  - integer "freq"

  - table "info"

  - table "patch"

**Example**

This example prints all information related to the first RT Channel for the first fixture in the selection:

```lua
return function()
    -- Get the index number for the first RT Channel for the first fixture in the current selection 
    local channelRTIndex = GetRTChannels(SelectionFirst())[1]
    -- Print an error message if returnd index is nil
    if channelRTIndex == nil then
        ErrPrintf("Please select a fixture and try again")
        return
    end
    -- Print all information about the RT Channel if it does not return nil
    local rtChannel = GetRTChannel(channelRTIndex)
    if rtChannel == nil then
        Printf("An RTChannel could not be found. Please try to selct a different fixture and try again.")
        return
    end
    Printf("================= RT CHANNEL =================")
    Printf("ui_index_first = " .. rtChannel["ui_index_first"])
    Printf("dmx_lowlight = " .. rtChannel["dmx_lowlight"])
    Printf("dmx_highlight = " .. rtChannel["dmx_highlight"])
    Printf("dmx_default = " .. rtChannel["dmx_default"])
    Printf("freq = " .. rtChannel["freq"])
    Printf("rt_index = " .. rtChannel["rt_index"])
    Printf("========== RELATED DMX CHANNEL DUMP ==========")
    rtChannel["dmx_channel"]:Dump() -- Handle for relevant DMX channel
    Printf("============ RELATED FIXTURE DUMP ============")
    rtChannel["fixture"]:Dump() -- Handle for relevant fixture
    Printf("========== RELATED SUBFIXTURE DUMP ===========")
    rtChannel["subfixture"]:Dump() -- Handle for relevant subfixture
    Printf("=================== INFO =====================")
    Printf("normed_phaser_time = " .. rtChannel["info"]["normed_phaser_time"])
    Printf("================ INFO FLAGS ==================")
    Printf("group_master = " .. rtChannel["info"]["flags"]["group_master"])
    Printf("additive_master = " .. rtChannel["info"]["flags"]["additive_master"])
    Printf("solo = " .. rtChannel["info"]["flags"]["solo"])
    Printf("highlight = " .. rtChannel["info"]["flags"]["highlight"])
    Printf("lowlight = " .. rtChannel["info"]["flags"]["lowlight"])
    Printf("=================== PATCH ====================")
    Printf("break = " .. rtChannel["patch"]["break"])
    Printf("coarse = " .. rtChannel["patch"]["coarse"])
    Printf("fine = " .. rtChannel["patch"]["fine"])
    Printf("ultra = " .. rtChannel["patch"]["ultra"])
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getrtchannel.html>


## GetAttributeByUIChannel

**Signature**

```
GetAttributeByUIChannel(integer:ui channel index): light_userdata:reference_to_attribute
```

**Help page title:** `GetAttributeByUIChannel(integer)`

**Description**

The **GetAttributeByUIChannel** Lua function returns the handle to an attribute based on a "UI Channel Index". The index number can be found in the Parameter List.

**Arguments**

- **Integer**:

            The integer is the UI Channel index number.

**Return**

- **Handle**:

            The handle to the attribute connected to the UI Channel.

**Example**

This example prints the "native" address to the first attribute of the first fixture in the current selection:

```lua
return function()
    -- Get a handle to the first fixture in the current selection
    local fixtureIndex = SelectionFirst()
    -- Get the UI Channel Index number for the first attribute for the fixture
    local channelIndex = GetUIChannelIndex(fixtureIndex,0)
    -- Print the native address for the attribute with the handle
    Printf("The native addr for the attribute is: %s",GetAttributeByUIChannel(channelIndex):AddrNative())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getattributebyuichannel.html>


## FirstDmxModeFixture

**Signature**

```
FirstDmxModeFixture(light_userdata:dmxmode): light_userdata:fixture
```

**Help page title:** `FirstDmxModeFixture(handle)`

**Description**

The **FirstDmxModeFixture** Lua function returns a handle to the first fixture matching the supplied DMX mode.

**Arguments**

- **Handle**:

            This must be a handle to a DMX mode.

**Return**

- **Handle**:

            The returned handle to the first fixture matching the DMX mode.

**Example**

If it exists, this example prints the data connected to the first "Dimmer" fixture using "Mode 0" - if the fixture type exists in the show. It uses the Dump() functions:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function ()
    -- Get the handle for the Dimmer fixture.
    local fixtureTypeHandle = Patch().FixtureTypes['Dimmer']

    -- Check if fixture type returned something and provide feedback.
    if fixtureTypeHandle == nil then
        ErrPrintf("The fixture type does not exist in this show. Try adding it or edit this plugin.")
    else
        -- Get the handle for the DMX mode of a Dimmer fixture.
        local fixtureDMXMode = fixtureTypeHandle.DMXModes["Mode 0"]

        -- Check if fixtureDMXMode returned something and provide feedback.
        if fixtureDMXMode == nil then
            ErrPrintf("The fixture type does not contain a 'Mode 0' DMX mode. Try adding it or edit this plugin.")
        else
            -- Dumps information about the first fixture matching the DMX mode.
            Printf("=============== START OF DUMP ===============")
            FirstDmxModeFixture(fixtureDMXMode):Dump()
            Printf("================ END OF DUMP ================")
        end
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_firstdmxmodefixture.html>


## NextDmxModeFixture

**Signature**

```
NextDmxModeFixture(light_userdata:fixture): light_userdata:fixture
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetAttributeIndex

**Signature**

```
GetAttributeIndex(string:attribute_name): integer:attribute_index
```

**Help page title:** `GetAttributeIndex(string)`

**Description**

The **GetAttributeIndex **Lua function returns the (0 based) index number of the attribute definition based on the system name of the attribute.

**Arguments**

- **String**:

            The string text of the attribute system name.

**Return**

- **Integer**:

            The returned integer number represents the total amount of patched fixtures on all the stages in the show file.

**Example**

This example prints the index number of the attribute in Command Line History if it exists:

```lua
return function()
    -- store the returned index or nil of "Gobo1"
    local attributeIndex = GetAttributeIndex("Gobo1")
    -- Check if the returned value is not nil and print a useful feedback
    if attributeIndex~=nil then
        Printf("Attribute is index number %i", attributeIndex)
    else
        Printf("The attribute is not found")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getattributeindex.html>


## GetUIChannelIndex

**Signature**

```
GetUIChannelIndex(integer:subfixture_index, integer:attribute_index): integer:ui_channel_index
```

**Help page title:** `GetUIChannelIndex(integer, integer)`

**Description**

The **GetUIChannelIndex** Lua function returns the index integer matching a UI channel based on two index inputs.

**Arguments**

- **Integer**:

 The first integer is the patch index of a fixture.

- **Integer**:

 This integer is an attribute index (0-based). This can be found in the Attribute Definitions or by the GetAttributeIndex() Lua function.

Attribute Definition

Attributes are the building blocks of fixture types. The same building blocks are used throughout the console and they are what is controlled using the Encoder bar when operating fixtures.

Attributes definitions describe the relation between Main Attributes and sub-attributes.

Learn more in the Attribute Definition topic.

GetAttributeIndex()

The **GetAttributeIndex **Lua function returns the (0 based) index number of the attribute definition based on the system name of the attribute.

Learn more in the GetAttributeIndex() topic.

**Return**

- **Integer**:

 The returned integer to a channel function.

**Example**

This example prints the UI channel index of the "Dimmer" attribute of the first fixture in the current selection:

```lua
return function()
    -- Get the Attribute index and UIChannel indexes
    local attributeIndex = GetAttributeIndex("Dimmer")
    local uiChannelIndex = GetUIChannelIndex(SelectionFirst(),attributeIndex)
    -- End the function if any of the index return nil
    if (attributeIndex == nil or uiChannelIndex == nil) then
        ErrPrintf("Something went wrong, maybe your first selected fixture don't have a Dimmer - Please try again")
        return
    end
    Printf("The UI Channel Index is " .. uiChannelIndex)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getuichannelindex.html>


## GetChannelFunctionIndex

**Signature**

```
GetChannelFunctionIndex(integer:ui_channel_index, integer:attribute_index): integer:channel_function_index
```

**Help page title:** `GetChannelFunctionIndex()`

**Description**

The **GetChannelFunctionIndex** Lua function returns the integer matching a channel function based on two index inputs.

**Arguments**

- **Integer**:

 The first integer is a UI Channel Index. This can be found in the Parameter List or by the GetUIChannelIndex() Lua function.

- **Integer**:

 This integer is an Attribute Index (0-based). This can be found in the Attribute Definitions or by the GetAttributeIndex() Lua function.

**Return**

- **Integer**:

 The returned integer to a channel function.

**Example**

This example prints the indexes based on the fixture selection and the "Dimmer" attribute.

```lua
return function()
    -- Get the Attribute index and UIChannel index.
    local attributeIndex = GetAttributeIndex("Dimmer")
    local uiChannelIndex = GetUIChannelIndex(SelectionFirst(),attributeIndex)
    -- End the function if any of the index return nil.
    if (attributeIndex == nil or uiChannelIndex == nil) then
        ErrPrintf("Something wrong happened, maybe your first selected fixture don't have a Dimmer - Please try again")
        return
    end
    -- Get the Channel Function Index and store it in a variable.
    local channelFunctionIndex = GetChannelFunctionIndex(uiChannelIndex,attributeIndex)
    Printf("The UIChannel Index is: %i. The Attribute Index is: %i. The Channel Function Index is: %i", uiChannelIndex, attributeIndex, channelFunctionIndex)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getchannelfunctionindex.html>


## GetChannelFunction

**Signature**

```
GetChannelFunction(integer:ui_channel_index, integer:attribute_index): light_userdata:handle
```

**Help page title:** `GetChannelFunction(integer, integer)`

**Description**

The **GetChannelFunction** Lua function returns a handle to a channel function based on two index inputs.

**Arguments**

- **Integer**:

            The first integer is a UI Channel Index. This can be found in the Parameter List or by the GetUIChannelIndex() Lua function.

- **Integer**:

            This integer is an Attribute Index (0-based). This can be found in the Attribute Definitions or by the GetAttributeIndex() Lua function.

**Return**

- **Handle**:

            The returned handle to the channel function.

**Example**

This example prints the data connected to the handle. It uses the Dump() function:

  Dump()

  The **Dump()** function returns a string with information about the object, for instance, the name, class, path of the object, its properties, and children.

  Learn more in the Dump() topic.

```lua
return function()
    -- Select the first fixture in the current selection.
    local subfixtureIndex = SelectionFirst()
    -- End the function if there is no selection.
    if subfixtureIndex == nil then
        ErrPrintf("Please select a fixture with a Dimmer")
        return
    end
    -- Get the Attribute index and UIChannel index.
    local attributeIndex = GetAttributeIndex("Dimmer")
    local uiChannelIndex = GetUIChannelIndex(subfixtureIndex,attributeIndex)
    Printf("The UIChannel Index is: %i. The Attribute Index is: %i. ",uiChannelIndex, attributeIndex)
    -- End the function if any of the index return nil.
    if (attributeIndex == nil or uiChannelIndex == nil) then
        ErrPrintf("Something wrong happened, maybe your first selected fixture don't have a Dimmer - Please try again")
        return
    end
    -- The following prints the dump for the dimmer channel function.
    Printf("=============== START OF DUMP ===============")
    GetChannelFunction(uiChannelIndex,attributeIndex):Dump()
    Printf("================ END OF DUMP ================")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getchannelfunction.html>


## GetTokenName

**Signature**

```
GetTokenName(string:short_name): string:full_name
```

**Help page title:** `GetTokenName(string)`

**Description**

The **GetTokenName** Lua function returns a string with the full keyword based on the short version string input or nil if there is no corresponding keyword.

**Arguments**

- **String**:

 The string input should correspond to a short version of a keyword.

**Return**

- **String**:

 A string with the full keyword is returned.

- OR -

- **Nil**:

 If there is no corresponding keyword, then nil is returned.

**Example**

This example returns the full keyword matching the short "seq" string:

```lua
return function()
    -- Store a short string to be used as input
    local shortToken = 'seq'
    -- Get the full token name
    local tokenName = GetTokenName(shortToken)
    -- Print useful output if nil is not returned
    if tokenName ~= nil then
        Printf("The full version of '".. shortToken .. "' is '" .. tokenName .. "'")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_gettokenname.html>


## GetTokenNameByIndex

**Signature**

```
GetTokenNameByIndex(integer:token_index): string:full_name
```

**Help page title:** `GetTokenNameByIndex(int)`

**Description**

The **GetTokenNameByIndex** Lua function returns a string with the keyword based on the index number provided.

Each keyword is described in the Command Syntax and Keywords section.

**Arguments**

- **Integer**:

 The integer input is the index number for a corresponding keyword. There is no apparent logic to the index number and the keyword.

**Return**

- **String**:

 A string with the full keyword is returned.

- OR -

- **Nil**:

 If there is no corresponding keyword, then nil is returned.

**Example**

If the keyword exists, this example returns the keywords matching the first 443 index numbers:

```lua
return function()
    -- Create a variable to hold the keyword string
    local tokenName = ""
    -- Print the keywords to the first 443 indexes if possible
    for index = 1, 443, 1 do
        tokenName = GetTokenNameByIndex(index)
        if tokenName ~= nil then
            Printf("Token index " .. index .. " = " .. tokenName)
        end
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_gettokennamebyindex.html>


## SelectionNotifyBegin

**Signature**

```
SelectionNotifyBegin(light_userdata:associated_context): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SelectionNotifyObject

**Signature**

```
SelectionNotifyObject(light_userdata:object_to_notify_about): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SelectionNotifyEnd

**Signature**

```
SelectionNotifyEnd(light_userdata:associated_context): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)

---

[Back to index](README.md)
