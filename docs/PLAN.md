# Requirements
- Wrapper to call zoxide query, zoxide add
- Handling of timeouts, parsing
- Error handling and error reporting
- good code, error types, const strings, good documentation
- examples
- Intuitive calls. WithAll().WithScore()
- make go.doc have a brief overview 
- unit tested
- linter configured

# API
- Wrapper for 
 - `zoxide query -l -s <args..>`
 - `zoxide add <path> -s <score>`
 - `zoxide remove <path>`
 - `zoxide --version` (for debugging ?)
- Check implementation of other libraries and get ideas
- (Optional) Update score
- (Optional) Zoxide's Import/Export commands
- (Optional) Allow specifying binary path
- (Optional) Detection of env variables and DB path.
- (Optional) Linting and Unit tests in CICD


# Q/A
- Global package handles ? Are they bad ?
  - Yes.
- Mistakes in golang libraries
  - Too many global variables/functions
Poor error wrapping (not using fmt.Errorf with %w)
  - Exposing internal types in public APIs
- North star in terms of functionalities ?
  - Fix optional stuff
- Stateless execution or stateful ?
  - Stateful via client.
- How to allow specifying timeout without having it passed in each functions, or having a zoxide variable defined ?
  - yes. zoxide client
- Learnings from other wrapper libraries
  - netlink : `pkgHandle` is a mess
  
  - docker/client: Good - uses contexts everywhere, clean error types
