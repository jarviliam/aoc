cmake_minimum_required(VERSION 3.8)
project(
  Test
  DESCRIPTION "description"
  HOMEPAGE_URL https://github.com/jarviliam/aoc
  LANGUAGES CXX
)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

add_subdirectory(day1)
add_subdirectory(day2)
add_subdirectory(day3)
add_subdirectory(day4)
add_subdirectory(day5)
add_subdirectory(day6)
