cmake_minimum_required(VERSION 3.17)
project(csync_cmp_perf)

set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -lpthread -Wall -Wextra -Werror")

add_executable(csync_cmp_perf main.cpp)

find_package(benchmark REQUIRED)
add_executable(csync_cmp_perf_bench bench.cpp)
target_link_libraries(csync_cmp_perf_bench benchmark::benchmark)
