# 📒 worklog

We use this to track our work history, so in the future we can have a look.

## The format

We try to keep it simple

```
[priority] [finish date]/[time spent]: [name] [unique-id]; [title]; [description]; [tags]
```

With some simple rules:

- One record stays in one line. So `newline` character is the way we break them out.

Some examples:

In our case, we use this:

```
**- 25.01.2024/1: Sang Dang <sang@dang.to>; Prepare the Syva starter; Use the taijutsu template, add magic link login; starter,next.js,taijutsu,syva
```

Which can be translate like this:

> Sang Dang uses the taijutsu template, prepare the starter code for Syva project, it also has the magic link login. This task supposes to be done at 25.1.2024, will take 01 hour to do.
We put this task into starter, nextjs, taijutsu and syva categories. The task priority is medium.

## Glossary

### Priority segment

| Indicator  | Meaning |
| ------------- | ------------- |
| `*--`  | Low priority  |
| `**-`  | Medium priority  |
| `***`  | High priority  |

### Finish date

For now we stick with simple `DD.MM.YYYY` format.

### Time spent

We round up to the closest integer. Some special numbers might be used.

| Indicator  | Meaning |
| ------------- | ------------- |
| `1`  | 1 hour  |
| `2`  | 2 hours  |
| `0`  | Not work  |
| `999`  | Do not know  |

### Name

We use `Firstname Lastname` format. For example `Sang Dang`.

### Unique ID

It could be UUID, or user email. Do NOT mix!

### Title

It's a string, should be shorter than 60 characters. Accepts almost every characters, except: `:` and `;`.

### Description

It's also a string, which can be longer. But should write it short and simple. Accepts almost every characters, except: `:` and `;`.

### Tags

Used for sorting, filtering and categorize in the future. They have some simple rules:

- Case-insensitive, so `Tag-1` and `tag-1` is the same.
- Use `-` instead of spaces. `tag 1` will be translated to `tag-1`.
- Multiple tags separated by a comma, `,`. For example `tag-1, tag-2, tag-3`
- All spaces are stripped out. `tag-1,    tag-2` is the same as `tag-1,tag-2`

## Empty values

We simply leave it empty, for example the log below does not have `description` and `tags`:

```
**- 31.01.2024/1: Sang Dang <sang@dang.to>; Upgrade Nodejs;;
```

## Contributing

Feel free to open an issue to ask us anything. PR is welcome but we should discuss first to avoid wasting time.

