# Progress bars

Create, update and close progress bars shown in the grandMA3 UI.

## Functions

- [`StartProgress`](#startprogress)
- [`StopProgress`](#stopprogress)
- [`SetProgressText`](#setprogresstext)
- [`SetProgressRange`](#setprogressrange)
- [`SetProgress`](#setprogress)
- [`IncProgress`](#incprogress)

---


## StartProgress

**Signature**

```
StartProgress(string:name): integer:progressbar_index
```

**Help page title:** `StartProgress(string)`

**Description**

The StartProgress Lua function creates and displays a progress bar on all screens. A string input argument creates a title for the progress bar. The function returns a handle that is used to further interact with the progress bar.

Executing the function displays the progress bar on the screens. It only disappears using the StopProgress function - which needs the handle. So it is highly recommended to store the returned handle from the start function.

See the ProgressBar topic for more info regarding progress bars and links to other related functions.

**Arguments**

- **String**:

 The string is used as the title for the progress bar.

**Return**

- **Handle**:

 The returned handle is the identifier for the progress bar.

**Example**

This creates and displays a progress bar on all screens. The progress bar does not disappear using this example - see the example in the StopProgress (link above) function to remove:

```lua
return function()
    -- Create and display a progress bar with a title
    -- IMPORTANT: The Lua variable 'progressHandle' is needed to remove the progressbar again - StopProgress()
    progressHandle = StartProgress("ProgressBar Title")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_startprogress.html>


## StopProgress

**Signature**

```
StopProgress(integer:progressbar_index): nothing
```

**Help page title:** `StopProgress(handle)`

**Description**

The StopProgress Lua function removes a progress bar. A handle input argument defines which progress bar it removes. The progress bar must exist before it can be removed. Progress bars are created using the StartProgress function.

See the ProgressBar topic for more info regarding progress bars and links to other related functions.

**Arguments**

- **Handle**:

 The handle for the progress bar to be stopped.

**Return**

This function does not return anything.

**Example**

This example stops the progress bar created using the example in the StartProgress topic (link above):

```lua
return function()
    -- Stops and closes the progress bar with the matching handle
    StopProgress(progressHandle)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_stopprogress.html>


## SetProgressText

**Signature**

```
SetProgressText(integer:progressbar_index, string:text): nothing
```

**Help page title:** `SetProgressText(handle, string)`

**Description**

The SetProgressText Lua function defines a text string to be displayed in a progress bar next to the progress bar title text. The title cannot be changed after creation, but this text can be changed. It could be used to describe the current step in the progress.

A handle input argument defines which progress bar it defines a text for. The progress bar must exist for it to have a handle. Progress bars can be created using the StartProgress function.

See the ProgressBar topic for more info regarding progress bars and links to other related functions.

**Arguments**

- **Handle**:

 The handle for the progress bar.

- **String**:

 The text string to be displayed.

**Return**

This function does not return anything.

**Example**

This example sets a text string for the progress bar created using the example in the StartProgress topic (link above):

```lua
return function()
    -- Sets the text next to progress title
    SetProgressText(progressHandle, "- This is text next to the progress title")
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_setprogresstext.html>


## SetProgressRange

**Signature**

```
SetProgressRange(integer:progressbar_index, integer:start, integer:end): nothing
```

**Help page title:** `SetProgressRange(handle, integer, integer)`

**Description**

The SetProgressRange Lua function defines a range for a progress bar.

A handle input argument defines which progress bar it defines a range for. The progress bar must exist to have a handle. Progress bars can be created using the StartProgress function.

See the ProgressBar topic for more info regarding progress bars and links to other related functions.

**Arguments**

- **Handle**:

 The handle for the progress bar.

- **Integer**:

 The start value for the range.

- **Integer**:

 The end value for the range.

**Return**

This function does not return anything.

**Example**

This example sets a range for the progress bar created using the example in the StartProgress topic (link above):

```lua
return function()
    -- Sets the range of a progress bar with the matching handle
    SetProgressRange(progressHandle, 1, 10)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_setprogressrange.html>


## SetProgress

**Signature**

```
SetProgress(integer:progressbar_index, integer:value): nothing
```

**Help page title:** `SetProgress(handle, integer)`

**Description**

The SetProgress Lua function defines a value on the range for a progress bar. A handle input argument defines the progress bar. The progress bar needs have been created using the StartProgress function.

See the ProgressBar topic for more info regarding progress bars and links to other related functions.

**Arguments**

- **Handle**:

 The handle for the progress bar.

- **Integer**:

 The desired value indicating the current status or position of the progress bar.

**Return**

This function does not return anything.

**Example**

This example sets a range value for the progress bar created using the example in the StartProgress topic (link above):

```lua
return function()
    -- Sets the current value to 5 for a progress bar with the matching handle
    SetProgress(progressHandle, 5)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_setprogress.html>


## IncProgress

**Signature**

```
IncProgress(integer:progressbar_index[, integer:delta]): nothing
```

**Help page title:** `IncProgress(handle, integer)`

**Description**

The IncProgress Lua function changes the value on the range for a progress bar using an integer input. A handle input argument defines the progress bar. The progress bar needs to be created using the StartProgress() function.

StartProgress()

The StartProgress Lua function creates and displays a progress bar on all screens.

Learn more in the StartProgress() topic.

See the ProgressBar topic for more info regarding progress bars and links to other related functions.

**Arguments**

- **Handle**:

 The handle for the progress bar.

- **Integer**:

 The desired value for the range. This can be a negative value to decrease the value.

**Return**

This function does not return anything.

**Example**

These two examples increase and decrease the range value for the progress bar created using the example in the StartProgress topic (link above):

```lua
return function()
    -- Increase the current value for a progress bar with the matching handle.
    IncProgress(progressHandle, 1)
end
```

```lua
return function()
    -- Decrease the current value for a progress bar with the matching handle.
    IncProgress(progressHandle, -1)
end
```

Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_incprogress.html>

---

[Back to index](README.md)
