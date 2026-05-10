# grandMA3 Practical Plugin Guide

> Compiled: **2026-05-10**
> Targeted MA3 software version: **2.3.2.0** (released 2025-11-03, the most recent public release on the MA Lighting downloads page)
> Help URL base: `https://help.malighting.com/grandMA3/2.3/HTML/...`

## Version confirmation

The version target was confirmed by visiting the official MA Lighting downloads page, which currently lists **2.3.2.0** as the only "Software" entry (older 2.3.1.1 / 2.3.0.4 are in the Archive, no 2.4 entry exists).

- The community has been waiting for **2.4** since at least early 2026, and as of May 2026 no release date has been announced — it was teased at ISE 2026.
- Source: <https://www.malighting.com/downloads/products/grandma3/>
- Source (2.4 status, May 2026 forum posts): <https://forum.malighting.com/forum/thread/69779-version-2-4-release-date/>

So `help.malighting.com/grandMA3/2.3/HTML/...` is the current official help URL. Some sub-pages on `help.malighting.com` only exist under `2.0`/`2.1`/`2.2` because they have not changed; in that case the older URL is still authoritative for now (`Cmd`, `HookObjectChange`, `Plugin` keyword, etc., all redirect cleanly within the help system).

### What's likely to change in 2.4

Version-sensitive items in this guide that should be re-checked when 2.4 ships:

- The `Cmd` / `CmdIndirect` / `CmdIndirectWait` triple — MA staff have hinted at API tightening in forum posts.
- Object property names on `Cue` and `Appearance` (the `BackR`/`BackG`/`BackB`/`ImageR`/...`/ImageAlpha` set is post-2.0; field names have shifted across point releases).
- `package.path` defaults for `require` (changed in 2.0; could change again).
- OSC configuration UI fields under Menu > In & Out > OSC.
- Anything labelled "unverified — needs testing on the console" below.

---

## Topic index

| # | Topic | Description |
|---|---|---|
| 1 | [Plugin runtime model](01_plugin_runtime.md) | Plugin Pool objects, the `Plugin` keyword, `Main`/`Cleanup`/`Execute`, coroutines vs threads, `Timer`, LuaSocket, `require`, persistent state. |
| 2 | [`Cmd()` in depth](02_cmd_function.md) | `Cmd` (synchronous), `CmdIndirect` / `CmdIndirectWait`, undo handles, the cmdline syntax that works, quoting/escaping pitfalls. |
| 3 | [Sequence / Cue / Executor object model](03_sequence_cue_executor.md) | Cue iteration with the OffCue / CueZero gotcha, `cue.No / 1000` numbering, Appearance fields, active-cue detection, executor pages, firing cues. |
| 4 | [OSC inside MA3](04_osc_in_ma3.md) | OSC configuration UI, inbound built-in addresses + `/cmd`, outbound `SendOSC`, per-config destination, the legacy Node.js bridge convention. |
| 5 | [Object addressing and handles](05_addressing_handles.md) | `H#...` handles, `Addr` vs `AddrNative` vs `ToAddr`, `FromAddr`. |
| 6 | [Hooks & change notification](06_hooks_and_changes.md) | `HookObjectChange`, `Unhook` / `UnhookMultiple` / `DumpAllHooks`, why polling still wins for our bridge. |
| 7 | [Community plugins worth referencing](07_community_plugins.md) | Upstream `ArtGateOne/ma3apcmini`, `MA3_OSC_FEEDBACK`, MacTirney, hossimo, patopesto, imhofroger, Bambinito. |
| 8 | [Common gotchas](08_gotchas.md) | `Echo` vs `Printf`, monkey-patched helpers like `string.starts`, paths, locale, `HelpLua` ground truth, plugin reload, `Cmd` blocking. |

---

## Unverified items

Items flagged "unverified" in the source — checklist for future console testing. Each entry links to the topic file and section where the original wording lives verbatim.

- [ ] **Anything labelled "unverified — needs testing on the console" below.** (meta-anchor — see [Version confirmation › What's likely to change in 2.4](#whats-likely-to-change-in-24) above and the items below)
- [ ] **LuaSocket on the console.** *Console availability is unverified — needs testing on the console.* — [01 Plugin runtime model § 1.7 LuaSocket availability](01_plugin_runtime.md#17-luasocket-availability)
- [ ] **`PluginVars()` scope semantics.** *Unverified — needs testing for actual scope semantics.* — [01 Plugin runtime model § 1.9 Persistent state, `_G`, `PluginVars`, `UserVars`/`GlobalVars`](01_plugin_runtime.md#19-persistent-state-_g-pluginvars-uservarsglobalvars)
- [ ] **`Cmd` embedded-quote escaping.** *Unverified — escaping rules for embedded quotes inside `Cmd` strings differ from macros.* — [02 Cmd function § 2.5 Quoting and escaping pitfalls](02_cmd_function.md#25-quoting-and-escaping-pitfalls)
- [ ] **Appearance colour value range.** *Values are 0..255 floats (not 0..1) per the forum example. Unverified — confirm range on console; the help page is silent.* — [03 Sequence / Cue / Executor § 3.3 Cue Appearance — the property names that matter](03_sequence_cue_executor.md#33-cue-appearance--the-property-names-that-matter)
- [ ] **`GetCurrentCue()` while sequence is Off.** *Unverified — behaviour when the sequence is Off (only OffCue conceptually active).* — [03 Sequence / Cue / Executor § 3.4 Detecting the active cue at runtime](03_sequence_cue_executor.md#34-detecting-the-active-cue-at-runtime)
- [ ] **Distinguishing paused vs running.** *Distinguishing paused vs running from Lua is unverified — needs testing on the console.* — [03 Sequence / Cue / Executor § 3.5 Sequence playback state](03_sequence_cue_executor.md#35-sequence-playback-state)
- [ ] **Linux/console plugin path.** *Linux/console: under `/var/MALightingTechnology/...` — exact path unverified, needs testing on console.* — [08 Gotchas § 8.3 Path handling](08_gotchas.md#83-path-handling)

---

## Sources

Consolidated alphabetical bibliography (deduplicated) of every URL cited across the topic files.

### Official MA Lighting help (v2.3 unless noted)

- <https://help.malighting.com/grandMA3/2.0/HTML/lua_handle.html>
- <https://help.malighting.com/grandMA3/2.0/HTML/lua_objectfree.html#FromAddr>
- <https://help.malighting.com/grandMA3/2.0/HTML/lua_objectfree_cmdindirectwait.html>
- <https://help.malighting.com/grandMA3/2.0/HTML/lua_objectfree_setvar.html>
- <https://help.malighting.com/grandMA3/2.1/HTML/lua_object_toaddr.html>
- <https://help.malighting.com/grandMA3/2.1/HTML/lua_objectfree_createundo.html>
- <https://help.malighting.com/grandMA3/2.2/HTML/extended_command_line.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/command_syntax_keywords.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/cue_sequence.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/key_releasenotes.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/keyword_plugin.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/keyword_sendosc.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_hasactiveplayback.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_cmd.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getcurrentcue.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getexecutor.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_hookobjectchange.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/plugins.html>
- <https://help.malighting.com/grandMA3/2.3/HTML/remote_inputs_osc.html>
- <https://help2.malighting.com/Page/grandMA3/lua_Timer/en/1.6>
- <https://help2.malighting.com/Page/grandMA3/lua_addr/en/1.9>

### MA Lighting downloads / company

- <https://www.malighting.com/downloads/products/grandma3/>

### Forum threads

- <https://forum.malighting.com/forum/thread/3917-getting-user-global-vars-into-or-out-of-plugins/>
- <https://forum.malighting.com/forum/thread/4501-lua-plugin-executing/>
- <https://forum.malighting.com/forum/thread/67878-reading-appearance-color-data-in-lua/>
- <https://forum.malighting.com/forum/thread/68211-grandma3-vscode-workflow/>
- <https://forum.malighting.com/forum/thread/69779-version-2-4-release-date/>
- <https://forum.malighting.com/forum/thread/7593-find-currentcue/>
- <https://forum.malighting.com/forum/thread/7973-non-blocking-plugins/>
- <https://forum.malighting.com/forum/thread/7999-using-require-to-include-lua-modules/>
- <https://forum.malighting.com/forum/thread/8069-get-the-current-cue-number-of-selected-sequence/>
- <https://forum.malighting.com/forum/thread/8291-call-a-plugin-and-specify-more-than-one-argument/>
- <https://forum.malighting.com/forum/thread/8477-load-go-depending-on-if-sequence-is-running/>
- <https://forum.malighting.com/forum/thread/8567-get-state-active-inactive-and-type-of-executor-toggle-flash-temp-via-lua-script/>
- <https://forum.malighting.com/forum/thread/8744-lua-insert-cue-in-sequence/>
- <https://forum.malighting.com/forum/thread/9033-how-to-get-correct-plugin-path/>
- <https://forum.malighting.com/forum/thread/9151-find-executor-location-in-lua/>
- <https://forum.malighting.com/forum/thread/9223-control-ma3-onpc-with-akai-apcmini-mk2-osc-artgateone/>

### Community references

- <https://github.com/ArtGateOne/MA3_OSC_FEEDBACK>
- <https://github.com/ArtGateOne/ma3apcmini>
- <https://github.com/MacTirney/GrandMA3-API-Documentation>
- <https://github.com/hossimo/GMA3Plugins>
- <https://github.com/imhofroger/GMA3_LUA>
- <https://github.com/patopesto/GrandMA3-Plugins>
- <https://grandma3.bambinito.net/>
- <https://grandma3.bambinito.net/guides/plugin-intro/>
- <https://grandma3.bambinito.net/reference/v23/api/>
- <https://support.actentertainment.com/knowledgeBase/26283481>

---

## How this folder is organized

This folder was split out of a single monolithic markdown file (`grandMA3_practical_guide.md`) on 2026-05-10 to make individual topics browseable in isolation.

- Numeric prefixes (`01_`, `02_`, …) preserve the original reading order. You can read the files top-to-bottom for the full guide, or jump straight to a topic.
- Each topic file carries a `> Targets MA3 2.3.2.0` sub-header so it is not version-orphaned if read in isolation.
- The version pin lives at the top of this README — re-check anything in the **What's likely to change in 2.4** list when 2.4 ships.
- All "unverified" flags survive verbatim inside their topic file *and* are surfaced as a consolidated checklist in the [Unverified items](#unverified-items) section above. That checklist is the queue for future console testing.
- The [Sources](#sources) section is a deduplicated bibliography of every URL cited across the topic files; useful as a fast reference when you want to chase the primary citation for a particular claim.
