cmake_minimum_required(VERSION 3.8)
project(
  Test
  DESCRIPTION "description"
  HOMEPAGE_URL https://github.com/jarviliam/coa
  LANGUAGES CXX
)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

add_subdirectory(day1)
