# DMX / patch

Add fixtures, build multi-patches, find / check DMX or FID collisions, and read raw DMX values.

## Functions

- [`AddFixtures`](#addfixtures)
- [`CreateMultiPatch`](#createmultipatch)
- [`FindBestDMXPatchAddr`](#findbestdmxpatchaddr)
- [`CheckDMXCollision`](#checkdmxcollision)
- [`CheckFIDCollision`](#checkfidcollision)
- [`GetDMXValue`](#getdmxvalue)
- [`GetDMXUniverse`](#getdmxuniverse)

---


## AddFixtures

**Signature**

```
AddFixtures({'mode'=light_userdata:dmx_mode, 'amount'=integer:amount[, 'undo'=string:undo_text][, 'parent'=light_userdata:handle][, 'insert_index'=integer:value][, 'idtype'=string:idtype][, 'cid'=string:cid][, 'fid'=string:fid][, 'name'=string:name][, 'layer'=string:layer][, 'class'=string:class][, 'patch'={table 1..8: string:address}]}): boolean:success or nothing
```

**Help page title:** `AddFixtures(table)`

**Description**

The **AddFixture** Lua function adds fixtures to the patch. The argument for the function is a table, which must contain valid data for the function to succeed. The function returns a "true" boolean value if the addition was a success. The function must be run with the command line in the correct patch destination.

**Arguments**

- **Table**:

            The table must contain valid data. This is a list of possible table elements. It is not necessary to add all elements.

  - **mode**:

                This must be a **handle **to a valid "dmx_mode". This defines a specific fixture type in a specific mode.

  - **amount**:

                This is an **integer **number that defines how many fixtures should be added.

  - **name** (optional):

                This is a string with the name of the (first) fixture.

  - **fid** (optional):

                This is a **string** with the fixture's FID.

  - **cid **(optional):

                This is a **string** with the CID for the fixture. This table field is only valid if the "idtype" is not "Fixture".

  - **idtype** (optional):

                This is a **string** with the name of the ID Type. This is only needed if the type is different than "Fixture".

  - **patch** (optional):

                This is a **table** with up to eight **strings**. The string must indicate a universe and a start address in the universe. The two must be separated by a dot. Each table element is used for the up to eight DMX breaks in the patch.

  - **layer** (optional):

                This is a **string** with the layer name.

  - **class** (optional):

                This is a **string** with the class name.

  - **parent** (optional):

                This is a **handle** of the parent fixture. It is only needed if the fixture should be a sub-fixture of an existing fixture.

  - **insert_index** (optional):

                This is an **integer** indicating an insert index number.

  - **undo** (optional):

                This is a **string** with an undo text.

**Return**

- **Boolean** or **nil**:

            The function returns a true boolean if the AddFixture function succeeded. It does not return anything (nil) if it fails.

**Example**

This example adds a dimmer fixture with FID and CID 301 and patch address "10.001". It is a requirement that the generic dimmer type is already added to the show, that the ID and patch address are available, and that the stage is called "Stage 1". The example does not perform any check for availability.

```lua
return function()
    -- Change the command line destination to the root.
    Cmd("ChangeDestination Root")
    -- Enter the "Patch".
    Cmd('ChangeDestination "ShowData"."Patch"')
    -- Enter the fixture location for the "Stage 1" object.
    Cmd('ChangeDestination "Stages"."Stage 1"."Fixtures"')

    -- Create a table.
    local myAddFixtureTable = {}
    -- Set the mode to a 8-bit Dimmer fixture type. 
    myAddFixtureTable.mode = Patch().FixtureTypes.Dimmer.DMXModes["Mode 0"]
    -- Set the amount of fixtures.
    myAddFixtureTable.amount = 1
    -- Set the FID for the fixture.
    myAddFixtureTable.fid = "301"
    -- Set the IdType - it is not needed if the type is "Fixture".
    myAddFixtureTable.idtype = "Channel"
    -- Set the CID - Use only this when the "idtype" is different than Fixture.
    myAddFixtureTable.cid = "301"
    -- Set the name of the fixture.
    myAddFixtureTable.name = "AddedDimmer 301"
    -- Create a patch table with an address.
    myAddFixtureTable.patch = {"10.001"}

    -- Add the fixture to the patch using the table data. Store the result in a local variable.
    local success = AddFixtures(myAddFixtureTable)
    
    -- Provide some feedback.
    if success ~= nil then
        Printf("Fixture " .. myAddFixtureTable.fid .. " is added with patch address " .. myAddFixtureTable.patch[1])
    else
        Printf("AddFixture failed!")
    end
    
    -- Return the command line to the root destination.
    Cmd("ChangeDestination Root")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_addfixtures.html>


## CreateMultiPatch

**Signature**

```
CreateMultiPatch({light_userdata:fixture_handles}, integer:count[ ,string:undo_text]): integer:amount_of_multi-patch_fixtures_created
```

**Help page title:** `CreateMultiPatch({handles}, integer[, string])`

**Description**

The **CreateMultiPatch** Lua function creates a series of multi patch fixtures to a table of fixtures.

**Arguments**

- **Table**:

            The table must contain **handles **to the fixtures who should have the multi patch fixtures.

- **Integer**:

            The number of multi patch fixtures to create.

- **String** (optional):

            The string is an optional undo text. It needs to be in quotation marks.

**Return**

- **Integer** | **nil**:

            The returned integer indicates the amount of multi patch fixtures created.

**Example**

This example creates two multi patch fixtures to the first fixture (excluding the "Universal" fixture) in the patch.

```lua
return function()
    -- Enter Patch.
    Cmd("ChangeDestination Root");
    -- Enter the SetupPatch.
    Cmd("ChangeDestination 'ShowData'.'Patch'");

    -- Get the handle for the first fixture in the patch.
    local myFixture = Patch().Stages[1].Fixtures[2]
    -- Add the handle a list element in an table.
    local myFixtureTable = {myFixture}
    -- Add a variable with the amount of multipatch fixtures needed.
    local multiPatchAmount = 2

    -- Count the number of elements in the fixture table and store in a variable.
    local count = 0
    for _ in pairs(myFixtureTable) do
         count = count + 1
    end
    -- Create an unto text string.
    local undoText = string.format("Create %d multipatch fixtures for up to %d fixtures", multiPatchAmount, count)

    -- Create the multipatch fixtures to the each fixture handle in the table and store the returned value.
    local multiPatchAmount = CreateMultiPatch(myFixtureTable, multiPatchAmount, undoText)
    if multiPatchAmount ~= nil then
        Printf(multiPatchAmount .. " multi patch objects was created")
    else
        Printf("An error occured")
    end

    -- Return the command line destination to the root.
    Cmd("ChangeDestination Root")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_createmultipatch.html>


## FindBestDMXPatchAddr

**Signature**

```
FindBestDMXPatchAddr(light_userdata:patch, integer:starting_address, integer:footprint): integer:absolute_address
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## CheckDMXCollision

**Signature**

```
CheckDMXCollision(light_userdata:dmx_mode, string:dmx_address[ ,integer:count[ ,integer:break_index]]): boolean:no_collision_found
```

**Help page title:** `CheckDMXCollision(handle, string[, integer[, integer]])`

**Description**

The **CheckDMXCollision** Lua function checks if a specific DMX address range is available or already used.

It uses the number of DMX channels in a specific "DMX mode" of a fixture type to calculate the number of DMX channels that should be available from a specified DMX start address.

All fixture types have at least one defined DMX mode. But fixtures often have more than one mode. This Lua function uses a specific DMX mode of a fixture type.

**Arguments**

- **Handle**:

            The handle must be for a "DMX mode". This is used to calculate how many DMX channels should be available in the range.

- **String**:
This must be a DMX address expressed as a string. This defines the start of the range to be checked.

- **Integer **(optional)|**nil**:

            This optional integer is a count of subsequent "DMX Modes" that should also be checked. The default value is **1**.
For instance, if the provided "DMX Mode" uses 10 DMX channels and the count is set to 5, then there must be 50 unpatched DMX channels from the start address for a positive result.

- **Integer **(optional):

            This optional integer indicates the break_index. The default value is **0**, which is the first "DMX break" defined for the "DMX mode". All fixture types have at least one defined "DMX break".

**Return**

- **Boolean**:

            The function returns a boolean.

  - **True**:

                The DMX address is available as a start address.

  - **False**:

                The DMX address is unavailable as a start address for the calculated number of DMX channels.

**Example**

This example prints feedback to the DMX collision check based on a DMX address of "1.001" and the DMX mode of the first fixture in the current selection:

```lua
return function()
    -- Set the DMX universe - range 1-1024.
    local myDMXUniverse = 1
    -- Set the DMX address in the universe - range 1-512.
    local myDMXAddress = 1
    -- Set the optional count for the number of fixtures (break_index channel amount) to check.
    local myCount = 1
    -- Set the optional break_index number for fixtures with multiple breaks. 
    -- Default value is 0 to indicate the first break.
    local myBreakIndex = 0

    -- Creates the string used for the DMX address.
    local startOfRange = string.format("%d.%03d", myDMXUniverse, myDMXAddress)

    -- Check if there is a selection and exit if there isn't.
    if SelectionFirst() == nil then
        Printf("Please make a selection and try again.")
        return
    end
    -- This gets the handle for the first fixture a patched generic Dimmers 8-bit mode.
    local myDmxMode = GetSubfixture(SelectionFirst()).ModeDirect

    if myDmxMode == nil then
        -- Exit the function if the DMX mode returns nil.
       
    else
        -- Do the actual collision check and provide useful feedback.
        if CheckDMXCollision(myDmxMode, startOfRange, myCount, myBreakIndex) then
            Printf("The DMX address " .. startOfRange .. " is available.")
            return
        else
            Printf("The DMX address " .. startOfRange .. " cannot be used as a start address for this patch.")
            return
        end
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_checkdmxcollision.html>


## CheckFIDCollision

**Signature**

```
CheckFIDCollision(integer:fid[, integer:count[, integer:type]]): boolean:no_collision_found
```

**Help page title:** `CheckFIDCollision(integer[, integer[, integer]])`

**Description**

The **CheckFIDCollision** Lua function checks if a specific (range of) ID is available or already used. It can be used to check FID and any type of CID by adding a type integer.

**Arguments**

- **Integer**:

            The first integer is the ID that should be checked.

- **Integer **(optional):

            This optional integer is a count of subsequent IDs that should also be checked. The default value is **1**. For instance, if FID 21 to 25 should be checked, then the count integer should be **5**.

- **Integer **(optional):

            This optional integer indicates the IDType. The default value is **0**, which is the "Fixture" ID Type. See the example below for other valid integers.

**Return**

- **Boolean**:

            The function returns a boolean.

  - **True**:

                The ID is available.

  - **False**:

                The ID is already used.

**Example**

This example prints feedback to the FID check:

```lua
return function()
    -- Create a variable with the FID you want to check.
    local myFID = 2001
    -- Create a variable with the number of subsequent ID's to also check.
    local myCount = 10
    -- Create a variable with the IDType you want to check.
    -- Default value is 0. This is the "Fixture" type.
    -- Valid integers are:
    --- 0 = Fixture
    --- 1 = Channel
    --- 2 = Universal
    --- 3 = Houseligths (default name)
    --- 4 = NonDim (default name)
    --- 5 = Media (default name)
    --- 6 = Fog (default name)
    --- 7 = Effect (default name)
    --- 8 = Pyro (default name)
    --- 9 = MArker
    --- 10 = Multipatch
    local myType = 0

    -- Check if the count is more than one.
    if myCount > 1 then
        -- Check if there is a collision and print valid feedback.
        if CheckFIDCollision(myFID, myCount, myType) then
            Printf("The FID " .. myFID .. " to " .. (myFID + myCount) .. " is available.")
            return
        else
            Printf("The FID " .. myFID .. " to " .. (myFID + myCount) .. " gives an FID collision.")
            return
        end
    else
        if CheckFIDCollision(myFID, nil, myType) then
            Printf("The FID " .. myFID .. " is available.")
            return
        else
            Printf("The FID " .. myFID .. " gives an FID collision.")
            return
        end
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_checkfidcollision.html>


## GetDMXValue

**Signature**

```
GetDMXValue(integer:address[ ,integer:universe, boolean:mode_percent]): integer:dmx_value
```

**Help page title:** `GetDMXValue(integer[, integer, boolean])`

**Description**

The **GetDMXValue **Lua function returns a number indicating the DMX value of a specified DMX address.

**Arguments**

- **Integer**:

 The integer is the DMX address. This value should be from 1 to 512 if a universe integer is provided. If a universe is not provided, this should be the absolute DMX address ranging from 1 to 524 288.

- **Integer** (optional):

 The integer is the universe number.

- **Boolean** (optional):

 The boolean indicates if the returned value is in percent or DMX value.

  - True:

 The returned value is in percent. The range is 0 to 100.

  - False:

 The returned value is in DMX value. The range is 0 to 255.

**Return**

- **Integer** or **nil**:

 The returned integer value corresponds with the value of the selected DMX address or nil if the DMX address is not granted.

**Example**

This example prints the value for DMX address 2 in Universe 1 (if it is granted):

```lua
return function()
    -- This prints the value of DMX address 2 in universe 1 in a range of 0 to 255
    local address = 2       -- The DMX address
    local universe = 1      -- The DMX universe
    local percent = false   -- Readout in percent or DMX value
    local value = GetDMXValue(address, universe, percent)
    if value == nil then
        Printf("The DMX address did not return a valid value")
    else
        Printf("DMX address %i.%03d has a value of %03d", universe, address, value)
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getdmxvalue.html>


## GetDMXUniverse

**Signature**

```
GetDMXUniverse(integer:universe[ ,boolean:modePercent]): {integer:dmx_values}
```

**Help page title:** `GetDMXUniverse(integer[, boolean])`

**Description**

The **GetDMXUniverse **Lua function returns a table with the DMX channels and their current value.

**Arguments**

- **Integer**:

 The integer is the universe number. The valid range is 1 to 1024.

- **Boolean** (optional):

 The boolean indicates if the returned value is in percent or DMX value.

  - True:

 The returned value is in percent. The range is 0 to 100.

  - False:

 The returned value is in DMX value. The range is 0 to 255.

**Return**

- **Table**:

 The returned table lists all the DMX addresses and the corresponding values.

-- OR --

- **Nil**:

 Nil is returned if the universe is not granted or the input value is out of range.

**Example**

This example prints the table in a list for DMX universe 1 (if it is granted):

```lua
return function()
    -- This gets a table for universe 1 with the returned value in percent.
    local tableDMXUniverse = GetDMXUniverse(1,true)
    -- Check the returned table and print information if nil.
    if tableDMXUniverse == nil then
        Printf("No value is returned. The univer is not granted or input is out of range")
        return
    end
    -- Prints the table if not nil.
    for addr, value in ipairs(tableDMXUniverse) do
        Printf("DMX Addr: %i - DMX value : %i", addr, value)
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getdmxuniverse.html>

---

[Back to index](README.md)
