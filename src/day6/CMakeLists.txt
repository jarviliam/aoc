add_executable(day6 main.cpp)

add_custom_command(
  TARGET day6 POST_BUILD
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day6/input.txt"
  $<TARGET_FILE_DIR:day6>
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day6/test.txt"
  $<TARGET_FILE_DIR:day6>
)
