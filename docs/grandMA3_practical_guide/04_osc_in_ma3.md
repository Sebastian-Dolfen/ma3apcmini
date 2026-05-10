# 4. OSC inside MA3

> Targets MA3 2.3.2.0

This section is load-bearing for our project.

## 4.1 Configuration UI

Menu **>** In & Out **>** OSC. You configure up to several "OSC configurations" (numbered slots).

Per-configuration fields (from the help page):

| Field | Meaning |
|---|---|
| Preferred IP | Local interface |
| Port | Local port (in) and remote port (out) — note both are configured |
| Mode | UDP or TCP |
| Destination IP | Where outbound packets go |
| Prefix | Optional path prefix added to all addresses (no slashes inside it) |
| Enable Input / Output | Master enable per direction |
| Receive / Send | Toggle data flow |
| Receive Command / Send Command | Toggle the special `/cmd` channel |

- Source: <https://help.malighting.com/grandMA3/2.3/HTML/remote_inputs_osc.html>

## 4.2 Inbound (we send → MA3 receives)

Two ways to drive MA3 over OSC:

**A. Built-in addresses** (case-sensitive):

```
/Fader201          ,i,50            -- exec 201 to 50% (current page)
/Page1/Fader201    ,i,50            -- absolute, page 1
/Page1/Key201      ,si,Press,1      -- key down/up
/Page1/Knob201     ,if,1,0.5        -- exec knob
```

Pool-style addresses: Sequences, Masters, Groups, Presets, Sounds, Worlds, Plugins, Screen Configurations, and Timers can all be addressed by enumerated number. Argument types `si` (key press/release) and `sif` (fader 0..100%) are documented.

**B. The `/cmd` channel** — full command line over OSC:

```
/cmd ,s,Go+ Exec 402
/cmd ,s,FaderMaster Page 1.201 At 50 Fade 5
/cmd ,s,Off Sequence 5
```

This requires `Receive Command = Yes` on the OSC configuration. It is the most flexible option but trades latency for power (commands go through the parser).

- Source: <https://help.malighting.com/grandMA3/2.3/HTML/remote_inputs_osc.html>
- Source: <https://support.actentertainment.com/knowledgeBase/26283481>

## 4.3 Outbound (`SendOSC` from Lua/cmdline)

```
SendOSC <ConfigID> "/Address,<TypeTags>,<Value>[,<Value>...]"
```

- `ConfigID` (integer, 1-based) selects which OSC configuration row to send through. `1` = first configuration in the OSC menu.
- Type tags: `i` int32, `f` float32, `s` string, `b` blob, `T` true, `F` false, `N` null, `I` impulse, `t` timetag. `T`/`F`/`N`/`I`/`t` take no payload.
- Multiple values separated by additional commas.

Examples (verbatim from the help page):

```
SendOSC 1 "/Page1/Fader201,i,50"
SendOSC 1 "/Page1/Fader201,ii,100,5"
SendOSC 1 "/cmd,s,Store Cue 1"
```

Shortcut: `Sen` (in macros / cmdline).

- Source: <https://help.malighting.com/grandMA3/2.3/HTML/keyword_sendosc.html>

## 4.4 Per-call vs per-config destination

Outbound OSC always goes to the **Destination IP** of the selected configuration row — there is no per-call override. To send to a different host, configure a second row and use `SendOSC 2 ...`.

## 4.5 Reference: the legacy Node.js bridge convention

`ArtGateOne/ma3apcmini` (the project we forked) uses the template:

```
SendOSC %i "/%s%i,is,%i,%s"
```

— i.e. `/<wing-prefix><exec-number>,is,<state>,<colorhex>`. The Node.js side (`ma3apcminimk2color.js`) parses `/Fader101`, `/Key201`-style addresses for the inbound feedback. Useful for reverse-engineering existing scenes.

- Source: <https://github.com/ArtGateOne/ma3apcmini>
- Source: <https://forum.malighting.com/forum/thread/9223-control-ma3-onpc-with-akai-apcmini-mk2-osc-artgateone/>

---

[Back to README](README.md)
