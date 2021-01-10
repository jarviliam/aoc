add_executable(day10 main.cpp)

target_compile_features(day10 PRIVATE cxx_std_17)
add_custom_command(
  TARGET day10 POST_BUILD
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day10/input.txt"
  $<TARGET_FILE_DIR:day10>
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day10/test.txt"
  $<TARGET_FILE_DIR:day10>
)
