# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.17

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Disable VCS-based implicit rules.
% : %,v


# Disable VCS-based implicit rules.
% : RCS/%


# Disable VCS-based implicit rules.
% : RCS/%,v


# Disable VCS-based implicit rules.
% : SCCS/s.%


# Disable VCS-based implicit rules.
% : s.%


.SUFFIXES: .hpux_make_needs_suffix_list


# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /home/koykov/tools/ide/CLion/bin/cmake/linux/bin/cmake

# The command to remove a file.
RM = /home/koykov/tools/ide/CLion/bin/cmake/linux/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf/cmake-build-debug

# Include any dependencies generated for this target.
include CMakeFiles/csync_cmp_perf_bench.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/csync_cmp_perf_bench.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/csync_cmp_perf_bench.dir/flags.make

CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.o: CMakeFiles/csync_cmp_perf_bench.dir/flags.make
CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.o: ../bench.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.o -c /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf/bench.cpp

CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf/bench.cpp > CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.i

CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf/bench.cpp -o CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.s

# Object files for target csync_cmp_perf_bench
csync_cmp_perf_bench_OBJECTS = \
"CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.o"

# External object files for target csync_cmp_perf_bench
csync_cmp_perf_bench_EXTERNAL_OBJECTS =

csync_cmp_perf_bench: CMakeFiles/csync_cmp_perf_bench.dir/bench.cpp.o
csync_cmp_perf_bench: CMakeFiles/csync_cmp_perf_bench.dir/build.make
csync_cmp_perf_bench: /usr/lib64/libbenchmark.so.1.5.6
csync_cmp_perf_bench: /usr/lib64/librt.so
csync_cmp_perf_bench: CMakeFiles/csync_cmp_perf_bench.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking CXX executable csync_cmp_perf_bench"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/csync_cmp_perf_bench.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/csync_cmp_perf_bench.dir/build: csync_cmp_perf_bench

.PHONY : CMakeFiles/csync_cmp_perf_bench.dir/build

CMakeFiles/csync_cmp_perf_bench.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/csync_cmp_perf_bench.dir/cmake_clean.cmake
.PHONY : CMakeFiles/csync_cmp_perf_bench.dir/clean

CMakeFiles/csync_cmp_perf_bench.dir/depend:
	cd /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf/cmake-build-debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf/cmake-build-debug /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf/cmake-build-debug /home/koykov/go/src/github.com/koykov/lab/csync_cmp_perf/cmake-build-debug/CMakeFiles/csync_cmp_perf_bench.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/csync_cmp_perf_bench.dir/depend

