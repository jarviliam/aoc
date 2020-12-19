add_executable(day4 main.cpp)

add_custom_command(
  TARGET day4 POST_BUILD
  COMMAND ${CMAKE_COMMAND} -E copy_if_different "${PROJECT_SOURCE_DIR}/day4/input.txt"
  $<TARGET_FILE_DIR:day4>
)
