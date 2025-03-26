# 📒 worklog

Minimalist log standard

## The standard

```worklog
[priority] [finish date]/[time spent]: [name] [unique-id]; [title]; [description]; [tags]
```

- One record stays in one line. So `newline` character is the way we break them out.
- New record goes on top, push old records down. So we read the most recent log first.

## Examples

In our case, we use this:

```worklog
**- 25.01.2024/1: Sang Dang <sang@dang.to>; Prepare the Syva starter; Use the taijutsu template, add magic link login; starter,next.js,taijutsu,syva
```

Which can be translate like this:

> Sang Dang uses the taijutsu template, prepare the starter code for Syva project, it also has the magic link login. This task supposes to be done at 25.1.2024, will take 01 hour to do.
> We put this task into starter, nextjs, taijutsu and syva categories. The task priority is medium.

## Why?

### 🧠 Memory Retention

Our brains are prone to forgetting details of daily routines. Maintaining a work log helps in preserving a detailed account of tasks, ensuring nothing significant is overlooked.

#### 📈 Comprehensive Achievement Tracking

It allows for a comprehensive record of achievements and milestones. This not only helps in recognizing individual accomplishments but also aids in identifying patterns of success over time.

#### 🚀 Efficient Planning for Performance Reviews

When it comes to annual performance reviews, having a historical log provides tangible evidence of accomplishments. This, in turn, enables better self-assessment, clearer communication during reviews, and more informed goal-setting for the future.

## How to use

You can either use this template by:

- Click the **"Use this template"** button and follow the instruction
- Or using the script below:

```bash
pnpm dlx tiged websitesieutoc/worklog
```

Start editing the `WORKLOG` file.

That's it!

> [!TIP] > [tiged](https://github.com/tiged/tiged) is a community driven fork of degit, it makes copies of git repositories without history.

## Glossary

### Priority segment

| Indicator | Meaning         |
| --------- | --------------- |
| `*--`     | Low priority    |
| `**-`     | Medium priority |
| `***`     | High priority   |

> [!TIP]
> You are free to adjust it to fit your needs, for example you can make it become 5 levels, like `***--` if you want.

### Finish date

For now we stick with simple `DD.MM.YYYY` format.

### Time spent

We round up to the closest integer. Some special numbers might be used.

| Indicator | Meaning     |
| --------- | ----------- |
| `1`       | 1 hour      |
| `2`       | 2 hours     |
| `0`       | Not work    |
| `-`       | Do not know |

### Name

We use `Firstname Lastname` format. For example `Sang Dang`.

### Unique ID

The ID is for the worker identification. It could be UUID/CUID, or user email. Incremental integer also works. This is NOT the ID of the record.

> [!IMPORTANT]
> Do not mix! Use either email, or ID.

### Title

It's a string, should be shorter than 60 characters. Accepts almost every characters, except: `:` and `;`.

> [!TIP]
> It's best to use Title for project's name, or something can remind the main features you have developed.

### Description

It's also a string, which can be longer. But should write it short and simple. Accepts almost every characters, except: `:` and `;`.

> [!TIP]
> Use Description to keep a quick snapshot of how your works impact.
> For example "Improve 20% number of MAU" or "Reduce 50% of actions when using our app".

### Tags

Used for sorting, filtering and categorize in the future. They have some simple rules:

- Case-insensitive, so `Tag-1` and `tag-1` is the same.
- Use `-` instead of spaces. `tag 1` will be translated to `tag-1`.
- Multiple tags separated by a comma, `,`. For example `tag-1, tag-2, tag-3`
- All spaces are stripped out. `tag-1,    tag-2` is the same as `tag-1,tag-2`

## Empty values

We simply leave it empty, for example the log below does not have `description` and `tags`:

```worklog
**- 31.01.2024/1: Sang Dang <sang@dang.to>; Upgrade Nodejs;;
```

## Contributing

Feel free to open an issue to ask us anything. PR is welcome but we should discuss first to avoid wasting time.
