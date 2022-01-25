# Dotfiles
Personal dotfiles repo.

Why is it written in Golang? No particular reason - just wanted to test my skills.

- `pkg/dotfiles` - generic toolbox to declare a component.
- `pkg/registry` - personal components declarations.
- `resources/*` - personal files that are referenced in `pkg/registry` components.

## What I want to do next
- [ ] `require` / `trigger` semantics for dependency declarations between components.
- [ ] integration tests (e.g. with Vagrant)
- [ ] components versioning
