# Programmer / phaser / color

Drive the programmer's phaser layer, set programmer colour, read preset data, and trigger colorimeter measurements.

## Functions

- [`SetProgPhaser`](#setprogphaser)
- [`SetProgPhaserValue`](#setprogphaservalue)
- [`GetProgPhaser`](#getprogphaser)
- [`GetProgPhaserValue`](#getprogphaservalue)
- [`SetColor`](#setcolor)
- [`GetPresetData`](#getpresetdata)
- [`ColMeasureDeviceDarkCalibrate`](#colmeasuredevicedarkcalibrate)
- [`ColMeasureDeviceDoMeasurement`](#colmeasuredevicedomeasurement)

---


## SetProgPhaser

**Signature**

```
SetProgPhaser(integer:ui_channel_index, {['abs_preset'=light_userdata:handle], ['rel_preset'=light_userdata:handle], ['fade'=number:seconds], ['delay'=number:seconds], ['speed'=number:hz], ['phase'=number:degree], ['measure'=number:percent], ['gridpos'=integer:value], {['channel_function'=integer:value], ['absolute'=number:percent], ['absolute_value'=integer:value], ['relative'=number:percent], ['accel'=number:percent[, 'accel_type'=integer:enum_value(Enums.SplineType)]], ['decel'=number:percent[, 'decel_type'=integer:enum_value(Enums.SplineType)]], ['trans'=number:percent], ['width'=number:percent], ['integrated'=light_userdata:preset_handle]}}): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SetProgPhaserValue

**Signature**

```
SetProgPhaserValue(integer:ui_channel_index, integer:step, {['channel_function'=integer:value], ['absolute'=number:percent], ['absolute_value'=integer:value], ['relative'=number:percent], ['accel'=number:percent[, 'accel_type'=integer:enum_value(Enums.SplineType)]], ['decel'=number:percent[, 'decel_type'=integer:enum_value(Enums.SplineType)]], ['trans'=number:percent], ['width'=number:percent], ['integrated'=light_userdata:preset_handle]}): nothing
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetProgPhaser

**Signature**

```
GetProgPhaser(integer:ui_channel_index, boolean:phaser_only): {['abs_preset'=light_userdata:handle], ['rel_preset'=light_userdata:handle], ['fade'=float:seconds], ['delay'=float:seconds], ['speed'=float:hz], ['phase'=float:degree], ['measure'=float:percent], ['gridpos'=integer:value], ['mask_active_phaser'=integer:bitmask], ['mask_active_value'=integer:bitmask], ['mask_individual'=integer:bitmask], {['channel_function'=integer:value], ['absolute'=float:percent], ['absolute_value'=integer:value], ['relative'=float:percent], ['accel'=float:percent[, 'accel_type'=integer:enum_value(Enums.SplineType)]], ['decel'=float:percent[, 'decel_type'=integer:enum_value(Enums.SplineType)]], ['trans'=float:percent], ['width'=float:percent], ['integrated'=light_userdata:preset_handle]}}
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetProgPhaserValue

**Signature**

```
GetProgPhaserValue(integer:ui_channel_index, integer:step): {['channel_function'=integer:value], ['absolute'=number:percent], ['absolute_value'=number:value], ['relative'=number:percent], ['accel'=number:percent[, 'accel_type'=integer:enum_value(Enums.SplineType)]], ['decel'=number:percent[, 'decel_type'=integer:enum_value(Enums.SplineType)]], ['trans'=number:percent], ['width'=number:percent], ['integrated'=light_userdata:preset_handle]}
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## SetColor

**Signature**

```
SetColor(string:color_model('RGB', 'xyY', 'Lab', 'XYZ', 'HSB'), float:tripel1, float:tripel2, float:tripel3, float:brightness, float:quality, boolean:const_brightness): integer:flag
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## GetPresetData

**Signature**

```
GetPresetData(light_userdata:preset_handle[, boolean:phasers_only(default=false)[, boolean:by_fixtures(default=true)]]): table:phaser_data
```

**Help page title:** `GetPresetData(handle[, boolean[, boolean]])`

**Description**

The **GetPresetData** Lua function returns a table with the preset data based on the preset handle.

The returned table is quite complex and has tables inside the table.

**Arguments**

- **Handle**:

 The handle** **of the preset from which the data will** **be collected.

- **Boolean** | **nil** (optional):
This boolean determines whether the returned table should only contain phaser data. The default value is "false".

- **Boolean** (optional):

 This boolean defines if there should be an extra object in the returned table. The default value is "true". The extra table object has the key "by_fixtures", and it contains the same table content as the returned table, but the keys are the fixture ID number instead of the UI Channel Index.

**Return**

- **Table** | **nil**:

 The returned table contains the preset data. It has multiple levels of tables.

**Example**

This example prints information about the first level table in the preset data and the first level of the first fixture in the preset. It uses dimmer preset 1, which must exist.

```lua
return function()
    -- Get the handle for the first Dimmer preset.
    local myPreset = DataPool().PresetPools[1][1]
    -- Get the Preset Data of the handle.
    local myPresetData = GetPresetData(myPreset, false, false)
    -- Check if the GetPresetData returns something. 
    if myPresetData == nil then
        ErrPrintf("Dimmer preset 1 does not exist. Please create one and try again.")
        return
    end

    -- Print the myPresetData table.
    for Key, value in pairs(myPresetData) do
      if type(value) == "table" then
        Printf("Key: " .. Key .. " ; Value type is: " .. type(value))
      else
        Printf("Key: " .. Key .. " ; Value type is: " .. type(value) .. " ; Value: " .. value)
      end
    end

    -- Create a table object to hold all the integer keys in the myPresetData table.
    local myIntegerTableKeys = {}
    -- Fill the table.
    for key,_ in pairs(myPresetData) do
        if type(key) == "number" then
            table.insert(myIntegerTableKeys, key)
        end
    end
    -- Sort the table
    table.sort(myIntegerTableKeys)

    -- Print the elements of the fixture with the lowest ui_channel_index in the preset.
    local tableIndex = myIntegerTableKeys[1]
    if tableIndex ~= nil then
        Printf("=============== TABLE CONTENT START - Table Key: " .. tableIndex .." ===============")
        for Key, value in pairs(myPresetData[tableIndex]) do
          if type(value) == "table" then
            Printf("Key: " .. Key .. " ; Value type is: " .. type(value))
          else
            Printf("Key: " .. Key .. " ; Value type is: " .. type(value) .. " ; Value: " .. tostring(value))
          end
        end
        Printf("================ TABLE CONTENT END - Table Key: " .. tableIndex .." ================")
    end
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getpresetdata.html>


## ColMeasureDeviceDarkCalibrate

**Signature**

```
ColMeasureDeviceDarkCalibrate(nothing): integer:flag
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)


## ColMeasureDeviceDoMeasurement

**Signature**

```
ColMeasureDeviceDoMeasurement(nothing): table:values
```

_(no detail page available — function is documented in the signature dump but no `lua_objectfree_*.html` page exists on the public help site)_

Source: index <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (no per-function page)

---

[Back to index](README.md)
