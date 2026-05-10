# Host / build / version

Identify the host machine (OS, console subtype, serial), the running build/version, and the show-file status.

## Functions

- [`HostOS`](#hostos)
- [`HostType`](#hosttype)
- [`HostSubType`](#hostsubtype)
- [`HostRevision`](#hostrevision)
- [`SerialNumber`](#serialnumber)
- [`OverallDeviceCertificate`](#overalldevicecertificate)
- [`ReleaseType`](#releasetype)
- [`DevMode3d`](#devmode3d)
- [`Version`](#version)
- [`BuildDetails`](#builddetails)
- [`GetShowFileStatus`](#getshowfilestatus)
- [`ConfigTable`](#configtable)

---


## HostOS

**Signature**

```
HostOS(nothing): string:ostype
```

**Help page title:** `HostOS()`

**Description**

The **HostOS** Lua function returns a string with the type of operating system of the device where the plugin is executed (for instance, "Windows", "Linux", or "Mac").

**Arguments**

This function does not accept any arguments.

**Return**

- **String**:

 The returned string is the operating system of the grandMA3 hardware or grandMA3 onPC computer.

**Example**

This example prints the operating system of the device in the Command Line History:

```lua
return function()
    Printf("The HostOS is "..HostOS())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_hostos.html>


## HostType

**Signature**

```
HostType(nothing): string:hosttype
```

**Help page title:** `HostType()`

**Description**

The **HostType** Lua function returns a string with the host type of the device where the plugin is executed (for example, "Console" or "onPC").

**Arguments**

This function does not accept any arguments.

**Return**

- **String**:

 The returned string is the host type of the device.

**Example**

This example prints the host type of the device in the Command Line History:

```lua
return function()
    Printf("The HostType is "..HostType())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_hosttype.html>


## HostSubType

**Signature**

```
HostSubType(nothing): string:hostsubtype
```

**Help page title:** `HostSubType()`

**Description**

The **HostSubType** Lua function returns a string with the host sub type of the station where the plugin is executed (for example, "FullSize", "Light", "RPU", "onPCRackUnit", "Undefined").

**Arguments**

This function does not accept any arguments.

**Return**

- **String**:

 The returned string is the host sub-type of the device.

**Example**

This example prints the host sub-type of the station in the Command Line History:

```lua
return function()
    Printf("The HostSubType is "..HostSubType())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_hostsubtype.html>


## HostRevision

**Signature**

```
HostRevision(nothing): string:hostrevision
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SerialNumber

**Signature**

```
SerialNumber(nothing): string:serialnumber
```

**Help page title:** `SerialNumber()`

**Description**

The **SerialNumber **Lua function returns the serial number of the grandMA3 hardware or grandMA3 onPC.

**Arguments**

This function does not accept any arguments.

**Return**

- **String**:

 The returned string is the serial number of the grandMA3 hardware or grandMA3 onPC.

**Example**

This example prints the serial number in the Command Line History:

```lua
return function()
    Printf("Serial number: " .. SerialNumber())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_serialnumber.html>


## OverallDeviceCertificate

**Signature**

```
OverallDeviceCertificate(nothing): light_userdata:handle
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## ReleaseType

**Signature**

```
ReleaseType(nothing): string:releasetype
```

**Help page title:** `ReleaseType()`

**Description**

The ReleaseType Lua function returns a string with the type of release for the MA software. All the software versions available from MA Lighting will return "Release". Internally and during development, there can be other release types.

**Arguments**

This function does not accept any arguments.

**Return**

- **String**:

 The returned string is the release type of the grandMA3 software.

**Example**

This example prints the release type in the Command Line History:

```lua
return function()
    Printf("The ReleaseType is "..ReleaseType())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_releasetype.html>


## DevMode3d

**Signature**

```
DevMode3d(nothing): string:devmode3d
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## Version

**Signature**

```
Version(nothing): string:version
```

**Help page title:** `Version()`

**Description**

The Version Lua function returns the software version.

**Arguments**

This function does not accept any arguments.

**Return**

- **String**:

            The returned string is the version of the grandMA3 software.

**Example**

This example prints the software version in the Command Line History:

```lua
return function()
    Printf("Software version: %s", Version())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_version.html>


## BuildDetails

**Signature**

```
BuildDetails(nothing): table:build_details
```

**Help page title:** `BuildDetails()`

**Description**

The BuildDetails function returns a table with key-value pairs about the software build.

**Arguments**

This function does not accept any arguments.

**Return**

- **Build details**:

            This is the table with key-value pairs. These are the possible keys in the table:

  - **GitDate**: String with the date for the repository branch of the software.

  - **GitHead**: String with the branch of the repository.

  - **GitHash**: String with the hash for the repository.

  - **CompileDate**: String with the date for the compile.

  - **CompileTime**: String with the time for the compile.

  - **BigVersion**: String indicating the software version.

  - **SmallVersion**: String with the small version number of the software. Devices that only listen to a DMX data stream need to have this version to "understand" the streaming data.

  - **HostType**: String with the host type, for instance, "Console" or "onPC".

  - **HostSubType**: String with the host sub-type, for instance, "FullSize" or "Light".

  - **CodeType**: String showing the type of code, for instance, "Release".

  - **IsRelease**: Boolean indicating if the software is a release version.

**Example**

This example prints the content of the BuildDetails table:

```lua
return function()
    --Store the build detials table
    local myBuild = BuildDetails()
    --Print the content of the table
    Printf("GitDate: " .. myBuild.GitDate)
    Printf("GitHead: " .. myBuild.GitHead)
    Printf("GitHash: " .. myBuild.GitHash)
    Printf("CompileDate: " .. myBuild.CompileDate)
    Printf("CompileTime: " .. myBuild.CompileTime)
    Printf("BigVersion: " .. myBuild.BigVersion)
    Printf("SmallVersion: " .. myBuild.SmallVersion)
    Printf("HostType: " .. myBuild.HostType)
    Printf("HostSubType: " .. myBuild.HostSubType)
    Printf("CodeType: " .. myBuild.CodeType)
    Printf("IsRelease: " .. tostring(myBuild.IsRelease))
end
```

              Related Lua Functions

- Version()

- HostType()

- HostSubType()

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_builddetails.html>


## GetShowFileStatus

**Signature**

```
GetShowFileStatus(nothing): string:showfile_status
```

**Help page title:** `GetShowFileStatus()`

**Description**

The **GetShowFileStatus** Lua function returns a string with the current device's show file status, for example, "NoShow", "ShowLoaded", "ShowDownloaded", "ShowSaving", and "DataNegotiationActive".

**Arguments**

This function does not accept any arguments.

**Return**

- **String**:

 The returned string is the enum string from "Enums.ShowFileStatus" that matches the current status.

**Example**

This example prints the current device's show file status in the Command Line History:

```lua
return function ()
    -- Prints the current showfile status
    Printf("ShowfileStatus: "..GetShowFileStatus())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getshowfilestatus.html>


## ConfigTable

**Signature**

```
ConfigTable(nothing): table:config_details
```

**Help page title:** `ConfigTable()`

**Description**

The **ConfigTable** Lua function returns a table with some configuration information. This is information only. The function does not have any actual functions. The table is not sorted.

**Arguments**

This function does not accept any arguments.

**Return**

- **Table**:

 The returned table contains key value pairs with configuration information. See the example below.

**Example**

This example prints the content of the returned table.

```lua
return function ()
    -- Prints the content of the ConfigTable
    for key,value in pairs(ConfigTable()) do
        Printf(key .. " : " .. value)
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_configtable.html>

---

[Back to index](README.md)
