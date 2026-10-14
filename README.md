# PeerPlay

A watch party system. Several people join a room, and a movie plays for all of
them at the same moment. When one person hits pause, everyone pauses. When
someone joins twenty minutes late, their video jumps straight to where everyone
else is.

This repository currently contains **no code**. It contains the design thinking
for code you are going to write yourself.

---

## The one idea that governs everything

There are two completely separate problems hiding inside "watch a movie
together":

**Problem 1 — moving the video bytes.** Getting roughly 500 MB to 4 GB of video
data from a server into each person's browser. This is a solved problem. A web
server sends the file, the browser asks for the parts it needs, the video
element plays them. Your Go room server never touches a single byte of video.

**Problem 2 — agreeing on where the playhead should be.** Getting four browsers
to all show frame 178,240 at the same wall-clock moment. This moves a few
hundred bytes of JSON that say, in effect, "the movie is playing, and it was at
743.2 seconds as of server time T". This is the interesting problem, and it is
the one you are building.

Almost every confusing thing about watch party systems comes from mixing these
two up. They travel over different connections, they fail in different ways, and
they need different fixes. A person whose video is buffering has a Problem 1
failure. A person whose video is 4 seconds behind everyone else has a Problem 2
failure. The doc set keeps them apart from beginning to end.

---

## What these docs are

A learning path, not a specification. You have finished the Go language
fundamentals but have never written an HTTP server. The docs take you from that
point to a running, observable, multi-instance system, in eleven phases, over
however many months you want.

**There is no implementation code in any of these files, on purpose.** You will
find JSON message sketches, ASCII diagrams, and descriptions of what a type
holds. You will not find a Go function you can copy. Copying a solution to a
concurrency problem teaches you nothing; being stuck on one for two hours
teaches you a lot.

---

## Reading order

Read these two first, before anything else:

| File | What it gives you |
|---|---|
| [docs/01-prerequisites.md](docs/01-prerequisites.md) | What you must know before you start, what to learn while building, what to ignore for now |
| [docs/02-problem-model.md](docs/02-problem-model.md) | The system described in plain terms, before any architecture |

Then use [docs/00-roadmap.md](docs/00-roadmap.md) as your spine. It lists eleven
phases. Each phase tells you what to build and points at the deeper files when
you need them.

The remaining files are references. Open one when the roadmap sends you there:

| File | Topic |
|---|---|
| [docs/03-architecture-stages.md](docs/03-architecture-stages.md) | Four architectures, and the pressure that forces each next one |
| [docs/04-domain-model.md](docs/04-domain-model.md) | The eight concepts the system is made of |
| [docs/05-protocol.md](docs/05-protocol.md) | What the HTTP and WebSocket messages must carry |
| [docs/06-concurrency.md](docs/06-concurrency.md) | Every place two goroutines can collide, and how to decide what to do about it |
| [docs/07-synchronization.md](docs/07-synchronization.md) | Keeping four playheads in step across a network with unequal delays |
| [docs/08-reliability-failures.md](docs/08-reliability-failures.md) | Twelve ways this breaks, and what should happen for each |
| [docs/09-observability-sre.md](docs/09-observability-sre.md) | How you find out the system is sick before your users tell you |
| [docs/10-security-deployment.md](docs/10-security-deployment.md) | Locking it down, and getting it onto a real machine |
| [docs/11-thinking-like-an-architect.md](docs/11-thinking-like-an-architect.md) | The repeatable process, and how this project builds an SRE mindset |
| [docs/12-exercises.md](docs/12-exercises.md) | Decide-first design problems, and failure-injection drills you run on your laptop |

---

## What "done" looks like

Not a finished product. The measure is whether you can explain your own
decisions to another engineer.

By Phase 5 you should be able to say, without looking anything up, why the room
state lives inside a single goroutine instead of behind a mutex, and why the
client sends `{"type": "PAUSE"}` instead of `{"playing": false}`. Those two
answers are worth more than the running code.

---

## Assumed media source

The design assumes **you serve the video yourself** — an MP4 or an HLS playlist
from your own server or from object storage you control. That gives the browser
a real `<video>` element, which reports its position to the millisecond and
seeks to an exact time. It is also the version where you own both ends, which
makes it far better for learning.

[docs/02-problem-model.md](docs/02-problem-model.md) has a section on what
changes if you swap in a third-party embed such as the YouTube IFrame Player
API instead. Short version: the player becomes a black box, seeking gets coarse,
ads appear without warning, and your drift tolerance has to grow a lot.

Nothing here requires bypassing DRM, scraping a streaming service, or
redistributing content you do not have the rights to. Use your own files, public
domain films, or a source you are licensed to use.
