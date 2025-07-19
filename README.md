# Discord-Social-Go

This project is currently in development. And is **not** ready for use.

## Getting Started

To run this project you need to provide the [Discord-Social-SDK](https://discord.com/developers/social-sdk) by yourself.

By creating an application on the [Discord Developer Portal](https://discord.com/developers/applications), then going to `DISCORD SOCIAL SDK`. Entering your details, and then download the newest SDK for `C++`.

Then create a new folder called `lib` and these files:

- `/discord_social_sdk/include/cdiscord.h`
- `/discord_social_sdk/bin/release/discord_partner_sdk.dll`
- `/discord_social_sdk/lib/release/libdiscord_partner_sdk.dylib`
- `/discord_social_sdk/lib/release/libdiscord_partner_sdk.so`

in the `lib` folder.

## Testing

To test this project you have a couple of test program in the `/test` folder.

You always need to add the Flag `-appid=<int64>` for the Discord Application ID. Run the test with the `-help` flag to get all Flags.
