add_executable(day8 main.cpp)

add_custom_command(
  TARGET day8 POST_BUILD
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day8/input.txt"
  $<TARGET_FILE_DIR:day8>
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day8/test.txt"
  $<TARGET_FILE_DIR:day8>
)
