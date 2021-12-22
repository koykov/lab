#include <iostream>
#include <atomic>
#include <benchmark/benchmark.h>
#include <mutex>
#include <thread>

const int TN = 1000;

struct x {
    int time;
    int allow;
    int acc;
    int nurlTime;
};

struct ax {
    int time;
    int allow;
    std::atomic<int> acc;
    std::atomic<int> nurlTime;
};

void bench_atomic_part(benchmark::State& b) {
    while (b.KeepRunning()) {
        ax s{1, 2, {3}, {4}};

        std::vector<std::thread> ths;
        ths.reserve(TN);

        for (int i = 0; i < TN; i++) {
            ths.emplace_back(std::thread([&s]() {
                s.acc.store(5, std::memory_order_relaxed);
            }));
        }

        for (auto &th: ths) {
            th.join();
        }
    }
}

void bench_mutex(benchmark::State& b) {
    while (b.KeepRunning()) {
        std::mutex mtx;
        x s{1, 2, 3, 4};

        std::vector<std::thread> ths;
        ths.reserve(TN);

        for (int i = 0; i < TN; i++) {
            ths.emplace_back(std::thread([&s, &mtx]() {
                mtx.lock();
                s.acc++;
                mtx.unlock();
            }));
        }

        for (auto &th: ths) {
            th.join();
        }
    }
}

BENCHMARK(bench_atomic_part);
BENCHMARK(bench_mutex);

BENCHMARK_MAIN();
