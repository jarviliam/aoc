add_executable(Day2 main.cpp)
add_custom_command(TARGET Day2 POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_if_different
    "${PROJECT_SOURCE_DIR}/day2/input.txt"
    $<TARGET_FILE_DIR:Day2>)
