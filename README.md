_[work-in-progress]_

# go-api-user
User-handling API example in GO.

## TODO (must-have)
- review error handling
- write tests
- identify race conditions
- review architecture

### TODO (low priority)
- review and possibly implement rest of the TODOs left in code
- dockerize

---

after CR:

- load .envs at startup
- struct / factory
- merge 'datasource' and 'repository' packages
- move server.Run() to main packages since it's handling whole app's startup
- added error handling and logging:
  https://github.com/MadCzarls/go-api-user/compare/main...refactoring#diff-2873f79a86c0d8b3335cd7731b0ecf7dd4301eb19a82ef7a1cba7589b5252261R10
- pointer usage:
- https://github.com/MadCzarls/go-api-user/compare/main...refactoring#diff-033285b6244f029398539e903f72206ff9c3444063844585751d36d895d5f5efR20
- read about:
- https://github.com/MadCzarls/go-api-user/compare/main...refactoring#diff-78f42ba40d0f10b08c73b7e6bb8376f398e249c963cf549e89591d6b6826b9a4R46
- remember about returning named variable
- https://github.com/MadCzarls/go-api-user/compare/main...refactoring#diff-78f42ba40d0f10b08c73b7e6bb8376f398e249c963cf549e89591d6b6826b9a4R19