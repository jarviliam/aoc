add_executable(day5 main.cpp)

add_custom_command(
  TARGET day5 POST_BUILD
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day5/input.txt"
  $<TARGET_FILE_DIR:day5>
)
