const std = @import("std");
const stdout = std.io.getStdOut().writer();
var gpallocator = std.heap.GeneralPurposeAllocator(.{}){};
var allocator = &gpallocator.allocator;

// This code was written AFTER the Go iterations.

const asc_u64 = std.sort.asc(u64);

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

    var illegal_n: ?u64 = null;

    for (numbers) |n, i| {
        if (i < preamble_len) {
            preamble[i] = n;
            continue;
        }

        if (!is_valid_sum(&preamble, n)) {
            illegal_n = n;
            break;
        }

        // Shift the array backwards and replace the last element.
        std.mem.copy(u64, preamble[0..], preamble[1..]);
        preamble[preamble_len - 1] = n;
    }

    var slice = array_sum(numbers, illegal_n.?).?;
    var conti = numbers[slice[0]..slice[1]];
    std.sort.sort(u64, conti, {}, asc_u64);

    try stdout.print("{}\n", .{conti[0] + conti[conti.len - 1]});
}

fn array_sum(numbers: []u64, sum: u64) ?[2]usize {
    for (numbers) |_, i| {
        var current_sum = numbers[i];

        var j = i + 1;
        while (j < numbers.len) : (j += 1) {
            if (current_sum == sum) {
                return [2]usize{ i, j - 1 };
            }
            if (current_sum > sum) {
                break;
            }
            current_sum += numbers[j];
        }
    }

    return null;
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
