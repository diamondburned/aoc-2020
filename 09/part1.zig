const std = @import("std");
const stdout = std.io.getStdOut().writer();
var gpallocator = std.heap.GeneralPurposeAllocator(.{}){};
var allocator = &gpallocator.allocator;

// This code was written AFTER the Go iterations.

pub fn main() !void {
    const f = try std.fs.cwd().openFile("input", .{ .read = true });
    const b = try f.reader().readAllAlloc(allocator, 1024000);

    var line_iterator = std.mem.split(b, "\n");
    var numbers_list = std.ArrayList(u64).init(allocator);

    while (true) {
        var line = line_iterator.next() orelse break;
        var lnum = std.fmt.parseInt(u64, line, 10) catch continue;
        try numbers_list.append(lnum);
    }

    var numbers = numbers_list.toOwnedSlice();

    const preamble_len = 25;
    var preamble: [preamble_len]u64 = [1]u64{0} ** preamble_len;

    for (numbers) |n, i| {
        if (i < preamble_len) {
            preamble[i] = n;
            continue;
        }

        if (!is_valid_sum(&preamble, n)) {
            try stdout.print("{} is invalid.\n", .{n});
            break;
        }

        // Shift the array backwards and replace the last element.
        std.mem.copy(u64, preamble[0..], preamble[1..]);
        preamble[preamble_len - 1] = n;
    }
}

fn is_valid_sum(preamble: []u64, sum: u64) bool {
    for (preamble) |n| {
        for (preamble) |m| {
            if (n + m == sum) {
                return true;
            }
        }
    }
    return false;
}
