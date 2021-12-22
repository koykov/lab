#include <iostream>
#include <atomic>
#include <vector>
#include <thread>
#include <mutex>

const int TN = 1000;
const int SAMPLES = 1e6;

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

void bench_atomic_part() {
    ax s{1, 2, {3}, {4}};

    std::vector<std::thread> ths;
    ths.reserve(TN);

    for (int i = 0; i < TN; i++) {
        ths.emplace_back(std::thread([&s]() {
            for (int i = 0; i < SAMPLES; i++) {
                s.acc.store(5, std::memory_order_relaxed);
            }
        }));
    }

    for (auto &th: ths) {
        th.join();
    }
}

void bench_mutex() {
    std::mutex mtx;
    x s{1, 2, 3, 4};

    std::vector<std::thread> ths;
    ths.reserve(TN);

    for (int i = 0; i < TN; i++) {
        ths.emplace_back(std::thread([&s, &mtx]() {
            for (int i = 0; i < SAMPLES; i++) {
                mtx.lock();
                s.acc++;
                mtx.unlock();
            }
        }));
    }

    for (auto &th: ths) {
        th.join();
    }
}

typedef std::chrono::time_point<std::chrono::high_resolution_clock,std::chrono::nanoseconds> time_ns_t;

time_ns_t now_ns() {
    return std::chrono::high_resolution_clock::now();
}

long since_ns(time_ns_t t) {
    auto ns = now_ns();
    long diff_ns = std::chrono::duration_cast<std::chrono::nanoseconds>(ns-t).count();
    return diff_ns;
}


int main() {
    auto now = now_ns();
    bench_atomic_part();
    std::cout << "bench atomic part took " << since_ns(now) << " ns" << std::endl;

    now = now_ns();
    bench_mutex();
    std::cout << "      bench mutex took " << since_ns(now) << " ns" << std::endl;

    return 0;
}
