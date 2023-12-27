package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// panic("Please implement the WelcomeMessage() function")
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// panic("Please implement the AddBorder() function")
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// panic("Please implement the CleanupMessage() function")
	message := strings.Trim(oldMsg, "\n *")
	message = strings.Join(strings.Fields(message), " ")
	return message
}

/*
在 C++, Python 和 Rust 中，你可以使用以下方式来实现类似的功能：

C++:
在 C++ 中，没有内置的函数可以直接完成这个任务，你需要手动编写函数来模拟这个行为。你可能需要使用 istringstream 来分割字符串，然后用 ostringstream 来重新拼接。

#include <sstream>
#include <vector>
#include <algorithm>
#include <cctype>
#include <iostream>

std::string cleanupMessage(const std::string& message) {
    std::istringstream iss(message);
    std::ostringstream oss;
    std::string word;
    while (iss >> word) {
        if (!oss.str().empty()) {
            oss << ' ';
        }
        oss << word;
    }
    return oss.str();
}

int main() {
    std::string message = "    Hello,   world!  ";
    std::cout << cleanupMessage(message) << std::endl;
}


Python:
Python 有内置的字符串方法和标准库函数来简化这个过程。

def cleanup_message(message):
    return " ".join(message.split())

# Usage
message = "    Hello,   world!  "
print(cleanup_message(message))



Rust:
在 Rust 中，你可以使用标准库中的字符串方法来实现。


fn cleanup_message(message: &str) -> String {
    message.split_whitespace().collect::<Vec<&str>>().join(" ")
}

fn main() {
    let message = "    Hello,   world!  ";
    println!("{}", cleanup_message(message));
}
在所有这些语言中，我们首先将字符串拆分为单词，然后再用单个空格将它们重新连接起来。这样做可以删除开头和结尾的空白，以及单词之间的多余空白。

*/
