# Filesystem & paths

Resolve grandMA3 path aliases, list directories, copy / sync files, and import or export show data.

## Functions

- [`GetPath`](#getpath)
- [`GetPathType`](#getpathtype)
- [`GetPathOverrideFor`](#getpathoverridefor)
- [`GetPathSeparator`](#getpathseparator)
- [`FileExists`](#fileexists)
- [`CopyFile`](#copyfile)
- [`CreateDirectoryRecursive`](#createdirectoryrecursive)
- [`SyncFS`](#syncfs)
- [`DirList`](#dirlist)
- [`Export`](#export)
- [`Import`](#import)
- [`ExportJson`](#exportjson)
- [`ExportCSV`](#exportcsv)

---


## GetPath

**Signature**

```
GetPath(string:path_type or integer:path_type(Enums.PathType)[ ,boolean:create]): string:path
```

**Help page title:** `GetPath(string[, boolean] | integer)`

**Description**

The **GetPath** Lua function returns a string with the path of a grandMA3 folder.

The function has two possible argument types - use one of them with each function call.

**Arguments**

- **String**:

 A text string with the folder name.

- **Boolean **(optional with string):

 If this boolean is true, then the folder at the path is created if it does not exist.

 - OR -

- **Integer**:

 An integer identifying an index in the "Enum.PathType" table.

            **Restriction:**

Folder creation only works with string arguments.

**Return**

- **String**:

 The returned string is the *first found* full path related to the provided argument.

**Example**

This example prints the paths of the show folder on the system monitor twice. It demonstrates the two different input types:

```lua
return function()
    -- This prints a path based on a string input and it creates the folder if it does not exists.
    Printf("Path of show files (string) is: " .. GetPath("shows", true))
    -- This prints the path based on an integer. The integer is looked-up using the 'PathType' enum.
    Printf("Path of show files (integer) is: " .. GetPath(Enums.PathType.Showfiles))
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getpath.html>


## GetPathType

**Signature**

```
GetPathType(light_userdata:target_object[ ,integer:content_type (Enums.PathContentType)]): string:path_type_name
```

**Help page title:** `GetPathType(handle[, integer])`

**Description**

The **GetPathType** Lua function returns a string with a name for the path type. This function can be useful when importing objects.

**Arguments**

- **Handle**:
The handle should match the object type for which the path type is needed.
- **Integer** (optional):
  The optional integer can be used to specify if the returned string should match the user path type or the system path type. See the example below.
The Enums.PathContentType can be used, or just use **0** for the system path and **1** for the user path.

**Return**

- **String**:
  The returned string is the name of the path type.

**Example**

This example prints the path type name for the first macro object - if it exists:

```lua
return function ()
    -- Get a handle to the first Macro.
    local myMacro = DataPool().Macros[1]
    if myMacro == nil then
        ErrPrintf("An error occurred, possibly because the first macro does not exist.")
        ErrPrintf("Please create one and try again.")
        return
    end
    -- Get the user name of the path type.
    local myPathTypeNameUser = GetPathType(myMacro, Enums.PathContentType.User)
    if myPathTypeNameUser ~= nil then
        Printf("The user name of the path type is: " .. myPathTypeNameUser)
    else
        ErrPrintf("There was an error getting the path type.")
    end

    -- Get the system name of the path type.
    local myPathTypeNameSystem = GetPathType(myMacro, Enums.PathContentType.System)
    if myPathTypeNameSystem ~= nil then
        Printf("The system name of the path type is: " .. myPathTypeNameSystem)
    else
        ErrPrintf("There was an error getting the path type.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getpathtype.html>


## GetPathOverrideFor

**Signature**

```
GetPathOverrideFor(string:path_type or integer:path_type(Enums.PathType), string:path[ ,boolean:create]): string:overwritten_path
```

**Help page title:** `GetPathOverrideFor(string|integer, string[, boolean])`

**Description**

The **GetPathOverrideFor** Lua function delivers a string with the path of a grandMA3 folder. The function is relevant when the path should be on a removable drive connected to a console.

**Arguments**

- **String**:

 A text string with the folder name.

- **String**:
The base path in a string format.

- **Boolean **(optional with string):

 If this boolean is true, then the folder at the path is created if it does not exist.

 - OR -

- **Integer**:

 An integer identifying an index in the "Enum.PathType" table.

- **String**:
The base path in a string format.

- **Boolean**:
If this boolean is true, then the folder at the path is created if it does not exist.

**Return**

- **String**:

 The returned string is the *first found* full path related to the provided argument.

**Example**

This example prints the override path of the macro folder on the system monitor. It should be run on a console with a removable drive connected.

```lua
return function ()
    -- Set a path for the first removable media.
    -- Set the initial value to nil.
    local myBasePath = nil
    -- Itereate the drives and find the first 'Removeable' drive and store the path.
    for _, value in ipairs(Root().Temp.DriveCollect) do
        local driveType = value.drivetype
        if driveType == "Removeable" then
            myBasePath = value.path
            break
        end
    end
    -- If no removeable drive was found, then provide feedback and exit the function.
    if myBasePath == nil then
        ErrPrintf("No removeable drive could be found. Please insert one and try again")
        return
    end

    -- Get the integer for the UserMacros path type.
    local myPathType = Enums.PathType.UserMacros

    -- Gey the string for the path override.
    local myOverridePath = GetPathOverrideFor(myPathType, myBasePath)
    -- Print the returned string.
    Printf("The path is: " .. myOverridePath)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getpathoverridefor.html>


## GetPathSeparator

**Signature**

```
GetPathSeparator(nothing): string:seperator
```

**Help page title:** `GetPathSeparator()`

**Description**

The **GetPathSeparator** function returns a string with the path separator for the operating system.

**Arguments**

This function does not accept any arguments.

**Return**

- **String**:

 The string is a single character indicating the path separator based on the operating system.

**Example**

This example prints the path separator:

```lua
return function()
    --- This prints the path seperator. It is different between a Linux and macOS (/) and a Windows (\) operating system.
    Printf("The path seperator is " .. GetPathSeparator())
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getpathseparater.html>


## FileExists

**Signature**

```
FileExists(string:path): boolean:result
```

**Help page title:** `FileExists(string)`

**Description**

The FileExists Lua function checks if a file exists and returns a boolean with the result.

**Arguments**

- **String**:

 The string must include the path and filename for the file that should be checked.

**Return**

- **Boolean**:

  - True / 1: The file exists.

  - False / 0: The file does not exist.

**Example**

This example returns feedback for the first file in the show file folder. The example uses the GetPath() and DirList() functions.

The **GetPath** Lua function delivers a string with the path of a grandMA3 folder.

Learn more in the GetPath() topic.

The **DirList** Lua function returns a table of files at a specified path.

Learn more in the DirList() topic.

```lua
return function ()
    -- Get the path to the show files.
    local path = GetPath(Enums.PathType.Showfiles)
    -- Get a table of files at the path.
    local dirTable = DirList(path)
    -- Get the file name for the first file.
    local firstFile = dirTable[1]['name']
    -- Create a string with the path and filename.
    local filepath = string.format("%s%s%s", path, GetPathSeparator(), firstFile)

    -- Check if the file exist and return useful feedback.
    if FileExists(filepath) then
        Printf('The file "' .. firstFile .. '" exist at path "' .. path .. '"')
    else
        Printf('The file "' .. firstFile .. '" does not exist')
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_fileexists.html>


## CopyFile

**Signature**

```
CopyFile(string:source_path, string:destination_path): boolean:result
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## CreateDirectoryRecursive

**Signature**

```
CreateDirectoryRecursive(string:path): boolean:result
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SyncFS

**Signature**

```
SyncFS(nothing): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## DirList

**Signature**

```
DirList(string:path[ ,string:filter]): table of {name:string, size:int, time:int}
```

**Help page title:** `DirList(string[, string])`

**Description**

The **DirList** Lua function returns a table of files at a specified path. The returned list can be filtered using an optional filter argument.

**Arguments**

- **String**:

            The desired path in a string format.

- **String** (optional):

            The optional filter string. The * can be used as a wildcard in the string. See the example below.

**Return**

- **Table**:

            The returned table has elements of other tables. Each of these table elements has the following keys:

  -
                name: The name of the file. The value of name is returned as a string.

  -
                size: The size of the file in bytes. The value of size is returned as a number.

  -
                time: The timestamp for the file. The value of time is returned as a number.

**Example**

This example prints the show files in the showfile directory. It uses the GetPath() function.

            The GetPath Lua function delivers a string with the path of a grandMA3 folder. Learn more in the GetPath() topic.

```lua
return function ()
    -- Get the path to the show files.
    local path = GetPath(Enums.PathType.Showfiles)
    -- Make a filter to only list .show files.
    local filter = "*show"
    -- Use the DirList function to get a table of the files.
    local returnTable = DirList(path, filter)

    -- Print the information of the files in the returned table.
    for _, value in pairs(returnTable) do
        Printf(value['name'] .. " - Size: " .. value['size'] .. " bytes - Time: " .. os.date("%c", value['time']))
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_dirlist.html>


## Export

**Signature**

```
Export(string:file_name, table:export_data): boolean:success
```

**Help page title:** `Export(filename, export_data)`

**Description**

The object-free **Export** Lua function exports a Lua table in XML format.

This Lua function correlates with the Import Lua function.

There is a related object version of Export.

**Arguments**

- **Filename**:

 This is a string containing the file name of the exported file. It should contain the file name, including the entire path. See the example below.

- **Export_data**:

 This is the data that is going to be exported. It should be a table object.

**Return**

- **Boolean**:
This function returns a boolean.

  - **True**:
The export was a success

  - **False**:
The export failed.

**Example**

To export the build details table, create a plugin with this code:

```lua
return function()
    -- 'BuildDetails()' creates a table with information about the software build.
    local build = BuildDetails()
    --The path and filename is stored in a variable.
    local exportPath = GetPath(Enums.PathType.Library) .. "/BuildDetails.xml"
    --The actual export (in xml format) using the path and the table - the result boolean stored in a variable.
    local success = Export(exportPath, build)
    --Print feedback about the export path.
    if success then
        Printf("The export was stored at: " .. exportPath)
    else
        Printf("The export failed")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_export.html>


## Import

**Signature**

```
Import(string:file_name): table:content
```

**Help page title:** `Import(string)`

**Description**

The object-free **Import** Lua function imports a Lua table in XML format.

        This function correlates to the Export function.

**Arguments**

- **String**:

          This is a string containing the file name of the desired imported file. It should contain the file name, including the entire path. See the example below.

**Return**

- **Table**:

          This is the imported table.

**Example**

This example imports the table exported using the example in the Export() function topic - please run that example before running this example.

```lua
return function ()
    -- Get the path for the exported table.
    local importPath = GetPath(Enums.PathType.Library) .. "/BuildDetails.xml"
    -- Check if the file exist and print relevant feedback.
    if importPath == nil then
        -- File does not exist.
        ErrPrintf("The desired file does not exist. Please add it or adjust the requested file name.")
    else
        -- Import the table.
        local importedTable = Import(importPath)
        -- Check if the import returned something and print relevant feedback.
        if importedTable == nil then
            -- Import didn't return anything.
            ErrPrintf("The import failed.")
        else
        -- Print some of the table content.
        Printf("CompileDate: " .. importedTable.CompileDate)
        Printf("CompileTime: " .. importedTable.CompileTime)
        Printf("BigVersion: " .. importedTable.BigVersion)
        Printf("HostType: " .. importedTable.HostType)
        Printf("HostSubType: " .. importedTable.HostSubType)
        Printf("CodeType: " .. importedTable.CodeType)
        end
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_import.html>


## ExportJson

**Signature**

```
ExportJson(string:file_name, table:export_data): boolean:success
```

**Help page title:** `ExportJson(filename, export_data)`

**Description**

The object-free **ExportJson** Lua function exports a Lua table in JSON format.

            **Known Limitation:**

The JSON file might not be formatted in proper JSON format.

**Arguments**

- **Filename**:

 This is a string containing the file name of the exported file. It should contain the file name, including the entire path. See the example below.

- **Export_data**:

 This is the data that is going to be exported. It should be a table object.

**Return**

- **Boolean**:
This function returns a boolean.

  - **True**:
The export was a success.

  - **False**:
The export failed.

**Example**

To export the build details table, create a plugin with this code:

```lua
return function()
    -- 'BuildDetails()' creates a table with information about the software build.
    local build = BuildDetails()
    --The path and filename is stored in a variable.
    local exportPath = GetPath(Enums.PathType.Library) .. "/BuildDetails.json"
    --The actual export (in JSON format) using the path and the table - the result boolean stored in a variable.
    local success = ExportJson(exportPath, build)
    --Print feedback about the export path.
    if success then
        Printf("The export was stored at: " .. exportPath)
    else
        Printf("The export failed.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_exportjson.html>


## ExportCSV

**Signature**

```
ExportCSV(string:file_name, table:export_data): boolean:success
```

**Help page title:** `ExportCSV(filename, export_data)`

**Description**

The object-free **ExportCSV** Lua function exports a Lua table in CSV format.

            **Known Limitation:**

The output CSV file might not formatted correctly.

**Arguments**

- **Filename**:

 This is a string containing the file name of the exported file. It should contain the file name, including the entire path. See the example below.

- **Export_data**:

 This is the data that is going to be exported. It should be a table object.

**Return**

- **Boolean**:
This function returns a boolean.

  - **True**:
The export was a success.

  - **False**:
The export failed.

**Example**

To export the build details table, create a plugin with this code:

```lua
return function()
    -- 'BuildDetails()' creates a table with information about the software build.
    local build = BuildDetails()
    --The path and filename is stored in a variable.
    local exportPath = GetPath(Enums.PathType.Library) .. "/BuildDetails.csv"
    --The actual export (in csv format) using the path and the table - the result boolean stored in a variable.
    local success = ExportCSV(exportPath, build)
    --Print feedback about the export path.
    if success then
        Printf("The export was stored at: " .. exportPath)
    else
        Printf("The export failed.")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_exportcsv.html>

---

[Back to index](README.md)
