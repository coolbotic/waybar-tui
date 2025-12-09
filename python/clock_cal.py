#!/usr/bin/env python3
import datetime
import calendar
import json
import re

# -------------------
# Current time
# -------------------
now = datetime.datetime.now()

# -------------------
# Catppuccin Mocha colors
# -------------------
FG_TITLE = "#F5C2E7"       # pink for main title
FG_MONTH = "#ffffff"       # month names white
FG_DAY_NAME = "#F5C2E7"    # day headers pink
FG_WEEKEND = "#8CAAEE"     # weekends blue
FG_WEEKDAY = "#ffffff"      # weekdays white
FG_TODAY = "#E78A4E"       # current day orange/red
FG_MOON = "#A6E3A1"        # moon phase green
FG_LINE = "#CCCCCC"        # thin white/gray separator

FONT_SIZE = 14
MONTH_WIDTH = 20
COLUMNS = 3
SPACING = 3
DAYS_IN_WEEK = 7

# -------------------
# Utility
# -------------------
def strip_markup(s):
    return re.sub(r"<.*?>", "", s)

def format_day(d, i, m):
    if d == 0:
        return "  "
    elif d == now.day and m == now.month:
        return f"<span foreground='{FG_TODAY}'><b>{d:2}</b></span>"
    elif i == 0 or i == 6:
        return f"<span foreground='{FG_WEEKEND}'>{d:2}</span>"
    else:
        return f"<span foreground='{FG_WEEKDAY}'>{d:2}</span>"

def month_calendar(m, y):
    cal = calendar.Calendar(firstweekday=6)
    month_name = f"<span foreground='{FG_TODAY}'><b>{calendar.month_name[m]} {y}</b></span>" if m == now.month else f"<span foreground='{FG_MONTH}'>{calendar.month_name[m]} {y}</span>"
    headers = " ".join([f"<span foreground='{FG_DAY_NAME}'>{d}</span>" for d in ["Su","Mo","Tu","We","Th","Fr","Sa"]])
    lines = [month_name]
    # Thin line under month name spanning exactly Sunday → Saturday
    lines.append(f"<span foreground='{FG_LINE}'>{'─' * MONTH_WIDTH}</span>")
    lines.append(headers)
    for week in cal.monthdayscalendar(y, m):
        line = [format_day(d, i, m) for i, d in enumerate(week)]
        lines.append(" ".join(line))
    return lines

# -------------------
# Build 12 months grid
# -------------------
months = [month_calendar(m, now.year) for m in range(1, 13)]
tooltip_lines = []

# Title line
top_line = f"📅 {now.strftime('%A, %d %B %Y')}"
tooltip_lines.append(f"<span foreground='{FG_TITLE}'>{top_line}</span>")
tooltip_lines.append("")

# Calendar layout
for row in range(0, 12, COLUMNS):
    month_block_lines = [months[row + i] for i in range(COLUMNS)]
    max_lines = max(len(b) for b in month_block_lines)
    for l in range(max_lines):
        parts = []
        for b in month_block_lines:
            if l < len(b):
                stripped = len(strip_markup(b[l]))
                pad = MONTH_WIDTH - stripped
                parts.append(b[l] + " " * pad)
            else:
                parts.append(" " * MONTH_WIDTH)
        tooltip_lines.append(f"<span font='{FONT_SIZE}'>{(' ' * SPACING).join(parts)}</span>")
    tooltip_lines.append("")

# -------------------
# Waybar-safe Moon phase
# -------------------
def simple_moon_phase(date):
    year = date.year
    month = date.month
    day = date.day
    if month < 3:
        year -= 1
        month += 12
    month += 1
    c = 365.25 * year
    e = 30.6 * month
    jd = c + e + day - 694039.09
    jd /= 29.5305882
    phase_index = int(jd) % 4
    phases = ["New Moon", "First Quarter", "Full Moon", "Last Quarter"]
    return phases[phase_index]

moon_phase_name = simple_moon_phase(now.date())
moon_icons = {
    "New Moon": "🌑",
    "First Quarter": "🌓",
    "Full Moon": "🌕",
    "Last Quarter": "🌗"
}
moon_text = f"{moon_icons[moon_phase_name]} Current Moon phase: {moon_phase_name}"
tooltip_lines.append(f"<span foreground='{FG_MOON}' font='{FONT_SIZE}'>{moon_text}</span>")

# -------------------
# Clock text
# -------------------
text = f"<span foreground='{FG_TITLE}'>🕛 {now.strftime('%I:%M %p')}</span>"

# -------------------
# Output
# -------------------
tooltip_text = "\n".join(tooltip_lines)
output = {
    "text": text,
    "tooltip": f"<span>{tooltip_text}</span>",
    "markup": "pango"
}

print(json.dumps(output))
