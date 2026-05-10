# grandMA3 Lua Object-Free API Reference

This document is a per-function reference for the Object-Free Lua API exposed by grandMA3 onPC / consoles. It was assembled by crawling every `lua_objectfree_*.html` page linked from the official help index at <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html> (grandMA3 software version **2.3** / docs marked Version 2.2 in the breadcrumb of most pages).

Function signatures are quoted from the source-of-truth dump at `docs/grandMA3_lua_functions.txt`. Where the help page text disagrees with that dump the discrepancy is flagged. Each section ends with a `Source:` link to the canonical help page.

Functions tagged **Used by APC plugin for:** are the ones the APC mini bridge in this repository relies on (or is most likely to rely on). Everything else is included for completeness.

---

## How this folder is organized

The reference is split across one Markdown file per API category. Each category file contains
its own short table of contents, then every function as an `H2` heading with the verbatim help-page
text (signature, description, arguments, return, examples, source URL).

Functions whose entry in this repo is tagged **Used by APC plugin for:** are flagged with [APC] in
the indexes below so you can scan for the ones that matter to the APC mini bridge.

Where a function is documented in the signature dump (`docs/grandMA3_lua_functions.txt`) but has no
dedicated `lua_objectfree_*.html` page on the public help site, the entry is kept as a stub
(signature only) inside whichever category fits semantically.

---

## Categories

- [Command execution & feedback](command_execution.md) - Run command-line strings, get feedback, log to System Monitor, and call commands on remote stations. _(9 functions)_
- [Host / build / version](host_build_version.md) - Identify the host machine (OS, console subtype, serial), the running build/version, and the show-file status. _(12 functions)_
- [Object access (handles, addresses, root pools)](object_access.md) - Resolve top-level pools, the current selection / user / executor, and convert between handles, addresses and integers. _(43 functions)_
- [Selection, subfixtures & UI channels](selection_subfixtures_ui_channels.md) - Walk the current selection, enumerate subfixtures, and inspect UI / RT channels, attributes and channel functions. _(29 functions)_
- [Programmer / phaser / color](programmer_phaser_color.md) - Drive the programmer's phaser layer, set programmer colour, read preset data, and trigger colorimeter measurements. _(8 functions)_
- [DMX / patch](dmx_patch.md) - Add fixtures, build multi-patches, find / check DMX or FID collisions, and read raw DMX values. _(7 functions)_
- [Hooks, callbacks & timers](hooks_callbacks_timers.md) - Subscribe to object-change events, wait on changes, manage hooks, and schedule recurring Lua timers. _(6 functions)_
- [Variables (Global/User/Plugin/Addon)](variables.md) - Access and mutate the four variable scopes (global / per-user / per-plugin / per-addon). _(8 functions)_
- [Filesystem & paths](filesystem_paths.md) - Resolve grandMA3 path aliases, list directories, copy / sync files, and import or export show data. _(13 functions)_
- [UI: input, dialogs & overlays](ui_input_dialogs_overlays.md) - Show modal dialogs (TextInput / Confirm / MessageBox), inspect displays and focus, and manipulate overlays. _(23 functions)_
- [Input simulation (keyboard / mouse / touch)](input_simulation.md) - Synthesize keyboard, mouse and touch events at either screen or object scope, and read system time. _(7 functions)_
- [Progress bars](progress_bars.md) - Create, update and close progress bars shown in the grandMA3 UI. _(6 functions)_
- [Property / column lookup](property_column_lookup.md) - Resolve property and attribute names to the column ids used by ObjectList / table accessors. _(2 functions)_
- [USB / MIDI device LEDs & buttons](usb_midi_leds_buttons.md) - Drive LEDs on attached USB / MIDI surfaces and read their button state - the core of the APC mini bridge. _(2 functions)_
- [Undo / desk / show state](undo_desk_show_state.md) - Open / close undo lists and inspect desk lock, remote-call activity and the unsaved-show flag. _(5 functions)_
- [Diagnostics & introspection](diagnostics_introspection.md) - Dump the API descriptor, read text-screen lines, and sample frame-rate / per-frame timing for debugging. _(6 functions)_
- [Messaging (Lua queues)](messaging.md) - Open named Lua message queues and post messages between Lua tasks or stations. _(3 functions)_

---

## Alphabetical function index

Functions marked **[APC]** have a `Used by APC plugin for:` note in their entry.

- [`AddFixtures`](dmx_patch.md#addfixtures)
- [`AddonVars`](variables.md#addonvars)
- [`BuildDetails`](host_build_version.md#builddetails)
- [`CallRealtimeLockedProtected`](command_execution.md#callrealtimelockedprotected)
- [`ChannelTable`](selection_subfixtures_ui_channels.md#channeltable)
- [`CheckDMXCollision`](dmx_patch.md#checkdmxcollision)
- [`CheckFIDCollision`](dmx_patch.md#checkfidcollision)
- [`ClassExists`](object_access.md#classexists)
- [`CloseAllOverlays`](ui_input_dialogs_overlays.md#closealloverlays)
- [`CloseMessageQueue`](messaging.md#closemessagequeue)
- [`CloseUndo`](undo_desk_show_state.md#closeundo)
- [`Cmd`](command_execution.md#cmd) **[APC]**
- [`CmdIndirect`](command_execution.md#cmdindirect) **[APC]**
- [`CmdIndirectWait`](command_execution.md#cmdindirectwait) **[APC]**
- [`CmdObj`](object_access.md#cmdobj)
- [`ColMeasureDeviceDarkCalibrate`](programmer_phaser_color.md#colmeasuredevicedarkcalibrate)
- [`ColMeasureDeviceDoMeasurement`](programmer_phaser_color.md#colmeasuredevicedomeasurement)
- [`ConfigTable`](host_build_version.md#configtable)
- [`Confirm`](ui_input_dialogs_overlays.md#confirm)
- [`CopyFile`](filesystem_paths.md#copyfile)
- [`CreateDirectoryRecursive`](filesystem_paths.md#createdirectoryrecursive)
- [`CreateMultiPatch`](dmx_patch.md#createmultipatch)
- [`CreateUndo`](undo_desk_show_state.md#createundo)
- [`CurrentEnvironment`](object_access.md#currentenvironment)
- [`CurrentExecPage`](object_access.md#currentexecpage)
- [`CurrentProfile`](object_access.md#currentprofile)
- [`CurrentScreenConfig`](object_access.md#currentscreenconfig)
- [`CurrentUser`](object_access.md#currentuser)
- [`DataPool`](object_access.md#datapool) **[APC]**
- [`DefaultDisplayPositions`](object_access.md#defaultdisplaypositions)
- [`DelVar`](variables.md#delvar)
- [`DeskLocked`](undo_desk_show_state.md#desklocked)
- [`DeviceConfiguration`](object_access.md#deviceconfiguration)
- [`DevMode3d`](host_build_version.md#devmode3d)
- [`DirList`](filesystem_paths.md#dirlist)
- [`DrawPointer`](ui_input_dialogs_overlays.md#drawpointer)
- [`DumpAllHooks`](hooks_callbacks_timers.md#dumpallhooks)
- [`Echo`](command_execution.md#echo) **[APC]**
- [`ErrEcho`](command_execution.md#errecho)
- [`ErrPrintf`](command_execution.md#errprintf) **[APC]**
- [`Export`](filesystem_paths.md#export)
- [`ExportCSV`](filesystem_paths.md#exportcsv)
- [`ExportJson`](filesystem_paths.md#exportjson)
- [`FileExists`](filesystem_paths.md#fileexists)
- [`FindBestDMXPatchAddr`](dmx_patch.md#findbestdmxpatchaddr)
- [`FindBestFocus`](ui_input_dialogs_overlays.md#findbestfocus)
- [`FindNextFocus`](ui_input_dialogs_overlays.md#findnextfocus)
- [`FindTexture`](ui_input_dialogs_overlays.md#findtexture)
- [`FirstDmxModeFixture`](selection_subfixtures_ui_channels.md#firstdmxmodefixture)
- [`FixtureType`](object_access.md#fixturetype)
- [`FromAddr`](object_access.md#fromaddr) **[APC]**
- [`FSExtendedModeHasDots`](ui_input_dialogs_overlays.md#fsextendedmodehasdots)
- [`GetApiDescriptor`](diagnostics_introspection.md#getapidescriptor)
- [`GetAttributeByUIChannel`](selection_subfixtures_ui_channels.md#getattributebyuichannel)
- [`GetAttributeColumnId`](property_column_lookup.md#getattributecolumnid)
- [`GetAttributeCount`](selection_subfixtures_ui_channels.md#getattributecount)
- [`GetAttributeIndex`](selection_subfixtures_ui_channels.md#getattributeindex)
- [`GetBlockInput`](ui_input_dialogs_overlays.md#getblockinput)
- [`GetButton`](usb_midi_leds_buttons.md#getbutton)
- [`GetChannelFunction`](selection_subfixtures_ui_channels.md#getchannelfunction)
- [`GetChannelFunctionIndex`](selection_subfixtures_ui_channels.md#getchannelfunctionindex)
- [`GetClassDerivationLevel`](object_access.md#getclassderivationlevel)
- [`GetCurrentCue`](object_access.md#getcurrentcue) **[APC]**
- [`GetDebugFPS`](diagnostics_introspection.md#getdebugfps)
- [`GetDisplayByIndex`](ui_input_dialogs_overlays.md#getdisplaybyindex)
- [`GetDisplayCollect`](ui_input_dialogs_overlays.md#getdisplaycollect)
- [`GetDMXUniverse`](dmx_patch.md#getdmxuniverse)
- [`GetDMXValue`](dmx_patch.md#getdmxvalue)
- [`GetExecutor`](object_access.md#getexecutor) **[APC]**
- [`GetFocus`](ui_input_dialogs_overlays.md#getfocus)
- [`GetFocusDisplay`](ui_input_dialogs_overlays.md#getfocusdisplay)
- [`GetObjApiDescriptor`](diagnostics_introspection.md#getobjapidescriptor)
- [`GetObject`](object_access.md#getobject) **[APC]**
- [`GetPath`](filesystem_paths.md#getpath)
- [`GetPathOverrideFor`](filesystem_paths.md#getpathoverridefor)
- [`GetPathSeparator`](filesystem_paths.md#getpathseparator)
- [`GetPathType`](filesystem_paths.md#getpathtype)
- [`GetPresetData`](programmer_phaser_color.md#getpresetdata)
- [`GetProgPhaser`](programmer_phaser_color.md#getprogphaser)
- [`GetProgPhaserValue`](programmer_phaser_color.md#getprogphaservalue)
- [`GetPropertyColumnId`](property_column_lookup.md#getpropertycolumnid)
- [`GetRemoteVideoInfo`](ui_input_dialogs_overlays.md#getremotevideoinfo)
- [`GetRTChannel`](selection_subfixtures_ui_channels.md#getrtchannel)
- [`GetRTChannelCount`](selection_subfixtures_ui_channels.md#getrtchannelcount)
- [`GetRTChannels`](selection_subfixtures_ui_channels.md#getrtchannels)
- [`GetSample`](diagnostics_introspection.md#getsample)
- [`GetScreenContent`](ui_input_dialogs_overlays.md#getscreencontent)
- [`GetSelectedAttribute`](object_access.md#getselectedattribute)
- [`GetShowFileStatus`](host_build_version.md#getshowfilestatus)
- [`GetSubfixture`](selection_subfixtures_ui_channels.md#getsubfixture)
- [`GetSubfixtureCount`](selection_subfixtures_ui_channels.md#getsubfixturecount)
- [`GetTextScreenLine`](diagnostics_introspection.md#gettextscreenline)
- [`GetTextScreenLineCount`](diagnostics_introspection.md#gettextscreenlinecount)
- [`GetTokenName`](selection_subfixtures_ui_channels.md#gettokenname)
- [`GetTokenNameByIndex`](selection_subfixtures_ui_channels.md#gettokennamebyindex)
- [`GetTopModal`](ui_input_dialogs_overlays.md#gettopmodal)
- [`GetTopOverlay`](ui_input_dialogs_overlays.md#gettopoverlay)
- [`GetUIChannel`](selection_subfixtures_ui_channels.md#getuichannel)
- [`GetUIChannelCount`](selection_subfixtures_ui_channels.md#getuichannelcount)
- [`GetUIChannelIndex`](selection_subfixtures_ui_channels.md#getuichannelindex)
- [`GetUIChannels`](selection_subfixtures_ui_channels.md#getuichannels)
- [`GetUIObjectAtPosition`](ui_input_dialogs_overlays.md#getuiobjectatposition)
- [`GetVar`](variables.md#getvar) **[APC]**
- [`GetVarVersion`](variables.md#getvarversion)
- [`GlobalVars`](variables.md#globalvars) **[APC]**
- [`HandleToInt`](object_access.md#handletoint) **[APC]**
- [`HandleToStr`](object_access.md#handletostr)
- [`HookObjectChange`](hooks_callbacks_timers.md#hookobjectchange) **[APC]**
- [`HostOS`](host_build_version.md#hostos)
- [`HostRevision`](host_build_version.md#hostrevision)
- [`HostSubType`](host_build_version.md#hostsubtype)
- [`HostType`](host_build_version.md#hosttype)
- [`Import`](filesystem_paths.md#import)
- [`IncProgress`](progress_bars.md#incprogress)
- [`IntToHandle`](object_access.md#inttohandle) **[APC]**
- [`IsClassDerivedFrom`](object_access.md#isclassderivedfrom)
- [`IsObjectValid`](object_access.md#isobjectvalid) **[APC]**
- [`Keyboard`](input_simulation.md#keyboard)
- [`KeyboardObj`](input_simulation.md#keyboardobj)
- [`LoadExecConfig`](object_access.md#loadexecconfig)
- [`MasterPool`](object_access.md#masterpool) **[APC]**
- [`MessageBox`](ui_input_dialogs_overlays.md#messagebox)
- [`Mouse`](input_simulation.md#mouse)
- [`MouseObj`](input_simulation.md#mouseobj)
- [`NeedShowSave`](undo_desk_show_state.md#needshowsave)
- [`NextDmxModeFixture`](selection_subfixtures_ui_channels.md#nextdmxmodefixture)
- [`ObjectList`](object_access.md#objectlist) **[APC]**
- [`OpenMessageQueue`](messaging.md#openmessagequeue)
- [`OverallDeviceCertificate`](host_build_version.md#overalldevicecertificate)
- [`Patch`](object_access.md#patch)
- [`PluginVars`](variables.md#pluginvars) **[APC]**
- [`PopupInput`](ui_input_dialogs_overlays.md#popupinput)
- [`PrepareWaitObjectChange`](hooks_callbacks_timers.md#preparewaitobjectchange)
- [`Printf`](command_execution.md#printf) **[APC]**
- [`Programmer`](object_access.md#programmer)
- [`ProgrammerPart`](object_access.md#programmerpart)
- [`Pult`](object_access.md#pult)
- [`RefreshLibrary`](object_access.md#refreshlibrary)
- [`ReleaseType`](host_build_version.md#releasetype)
- [`RemoteCallRunning`](undo_desk_show_state.md#remotecallrunning)
- [`RemoteCommand`](command_execution.md#remotecommand) **[APC]**
- [`Root`](object_access.md#root) **[APC]**
- [`SaveExecConfig`](object_access.md#saveexecconfig)
- [`SelectedDrive`](object_access.md#selecteddrive)
- [`SelectedFeature`](object_access.md#selectedfeature)
- [`SelectedLayout`](object_access.md#selectedlayout)
- [`SelectedSequence`](object_access.md#selectedsequence) **[APC]**
- [`SelectedTimecode`](object_access.md#selectedtimecode)
- [`SelectedTimer`](object_access.md#selectedtimer)
- [`Selection`](object_access.md#selection)
- [`SelectionComponentX`](selection_subfixtures_ui_channels.md#selectioncomponentx)
- [`SelectionComponentY`](selection_subfixtures_ui_channels.md#selectioncomponenty)
- [`SelectionComponentZ`](selection_subfixtures_ui_channels.md#selectioncomponentz)
- [`SelectionCount`](selection_subfixtures_ui_channels.md#selectioncount)
- [`SelectionFirst`](selection_subfixtures_ui_channels.md#selectionfirst)
- [`SelectionNext`](selection_subfixtures_ui_channels.md#selectionnext)
- [`SelectionNotifyBegin`](selection_subfixtures_ui_channels.md#selectionnotifybegin)
- [`SelectionNotifyEnd`](selection_subfixtures_ui_channels.md#selectionnotifyend)
- [`SelectionNotifyObject`](selection_subfixtures_ui_channels.md#selectionnotifyobject)
- [`SelectionTable`](selection_subfixtures_ui_channels.md#selectiontable)
- [`SendLuaMessage`](messaging.md#sendluamessage) **[APC]**
- [`SerialNumber`](host_build_version.md#serialnumber)
- [`SetBlockInput`](ui_input_dialogs_overlays.md#setblockinput)
- [`SetColor`](programmer_phaser_color.md#setcolor)
- [`SetLED`](usb_midi_leds_buttons.md#setled)
- [`SetProgPhaser`](programmer_phaser_color.md#setprogphaser)
- [`SetProgPhaserValue`](programmer_phaser_color.md#setprogphaservalue)
- [`SetProgress`](progress_bars.md#setprogress)
- [`SetProgressRange`](progress_bars.md#setprogressrange)
- [`SetProgressText`](progress_bars.md#setprogresstext)
- [`SetVar`](variables.md#setvar) **[APC]**
- [`ShowData`](object_access.md#showdata)
- [`ShowSettings`](object_access.md#showsettings)
- [`StartProgress`](progress_bars.md#startprogress)
- [`StopProgress`](progress_bars.md#stopprogress)
- [`StrToHandle`](object_access.md#strtohandle)
- [`SyncFS`](filesystem_paths.md#syncfs)
- [`TextInput`](ui_input_dialogs_overlays.md#textinput)
- [`Time`](input_simulation.md#time)
- [`Timer`](hooks_callbacks_timers.md#timer) **[APC]**
- [`ToAddr`](object_access.md#toaddr) **[APC]**
- [`Touch`](input_simulation.md#touch)
- [`TouchObj`](input_simulation.md#touchobj)
- [`Unhook`](hooks_callbacks_timers.md#unhook) **[APC]**
- [`UnhookMultiple`](hooks_callbacks_timers.md#unhookmultiple)
- [`UserVars`](variables.md#uservars) **[APC]**
- [`Version`](host_build_version.md#version)
- [`WaitModal`](ui_input_dialogs_overlays.md#waitmodal)
- [`WaitObjectDelete`](ui_input_dialogs_overlays.md#waitobjectdelete)
