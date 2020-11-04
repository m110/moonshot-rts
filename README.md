# Moonshot RTS

This is a game in development for [Github Game Off 2020](https://itch.io/jam/game-off-2020). Made in Go with [ebiten](https://github.com/hajimehoshi/ebiten).

## Why Go?

While I usually use Unity for game jams, I decided to try writing a game in Go this time.

* It's fun to build a simple engine from scratch, even if it takes longer time.
* Since it's GitHub's Game Off, I want this game to be fully open source.
* While Unity is a complete platform, it feels bloated at times and I had a terrible experience using it together with git.
* I hope the simplicity of Go's syntax can make it easier to understand the project for anyone looking to get into Go or gamedev.

## Theme

* It's a moonshot, because it's my first attempt at an RTS. 😅
* Building an empire out of a small settlement is a moonshot as well.

## Engine

I use [ebiten](https://github.com/hajimehoshi/ebiten) library for drawing images and basic game loop.

I'm trying to build a practical [Entity-Component-System](https://en.wikipedia.org/wiki/Entity_component_system) pattern.
What I did so far is heavily inspired by [EngoEngine/ecs](https://github.com/EngoEngine/ecs). However, I took a bit
different approach with interfaces. This is still under heavy development, though. 🙂

The TL;DR version of ECS is:

* The entire game is built out of **Entities**, **Components**, and **Systems**.
* An **entity** is a container for **components** with a unique ID.
* A **component** is a data structure, describing some set of common values.
* A **system** is where the game logic lives for a particular area.

For example, there's `DrawingSystem` that draws sprites on the screen.
It requires an entity with `Drawable` (a sprite) and `WorldSpace` (a position) components.

## Game

* Right now, my idea is a mix between Age of Empires style RTS, and Civilzation resource tiles
* However, I started out focusing on the engine, so I'm not sure where this will go yet. 😀

## Assets

This will possibly change, but right now:

* [Medieval RTS from Kenney](https://kenney.nl/assets/medieval-rts)
* [UI Pack: RPG Expansion from Kenney](https://kenney.nl/assets/ui-pack-rpg-expansion)
* OpenSans font from Google

## Running

```
go run ./cmd/rts
```
