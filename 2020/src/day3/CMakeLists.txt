add_executable(day3 main.cpp)

add_custom_command(
  TARGET day3 POST_BUILD
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day3/input.txt"
  $<TARGET_FILE_DIR:day3>
)
