add_executable(day11 main.cpp)

target_compile_features(day11 PRIVATE cxx_std_17)
add_custom_command(
  TARGET day11 POST_BUILD
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day11/input.txt"
  $<TARGET_FILE_DIR:day11>
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day11/test.txt"
  $<TARGET_FILE_DIR:day11>
)
