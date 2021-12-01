add_executable(day7 main.cpp)

add_custom_command(
  TARGET day7 POST_BUILD
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day7/input.txt"
  $<TARGET_FILE_DIR:day7>
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day7/test.txt"
  $<TARGET_FILE_DIR:day7>
)
