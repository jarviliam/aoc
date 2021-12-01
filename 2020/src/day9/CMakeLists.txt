add_executable(day9 main.cpp)

target_compile_features(day9 PRIVATE cxx_std_17)
add_custom_command(
  TARGET day9 POST_BUILD
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day9/input.txt"
  $<TARGET_FILE_DIR:day9>
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day9/test.txt"
  $<TARGET_FILE_DIR:day9>
)
