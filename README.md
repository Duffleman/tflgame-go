# tflgame

This is the backend service that runs the TFLGame. A game where you are provided a selection of letters and you have to guess the TFL station name.

Any case where response is not defined in this document, you will get a 204 status code, it completed the requested action.

## Auth

To generate the EC256 keys needed for JWT signing, use the following two commands.

```bash
openssl ecparam -genkey -name prime256v1 -noout -out ec_private.pem
openssl ec -in ec_private.pem -pubout -out ec_public.pem
```

## Entities

### User tag

Each user will get a tag, a tag consists of [A-Z] minimum 2, maximum 5. You can duplicate tags as the system tracks numbers that are assigned for you. So if you request "KLM" as your tag, you will be assigned "KLM#001". Additionally you can provide a pin to protect your tag so that your name is not sullied on the leaderboard.

### Scoring

### Game options

#### `include_random_spaces`

This option will include random spaces into your prompts, so that it is slightly trickier to guess. `WHITECHAPEL` will turn to `W HTCH PL`, and the spaces added will be random each time.

#### `change_letter_order`

This option will swap up the letters. `WHITECHAPEL` will return to `ITWLECHPEHA`. Making it quite a bit harder to guess what is going on.

#### `reveal_word_length`

This option will mean that when you're given prompts, you're also told how long the final answer is going to be.

## API

### `create_user`

#### Request

```json
{
	"handle": "DFL",
	"pin": "014410"
}
```

#### Response

```json
{
	"id": "user_000000C0KhWzlS5SMpISfDx5IF3aN",
	"handle": "DFL",
	"numeric": "001",
	"token": "eyj..."
}
```

### `authenticate`

#### Request

```json
{
	"handle": "DFL",
	"numeric": "001",
	"pin": "014410"
}
```

#### Response

```json
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN",
	"token": "eyj..."
}
```

Although the `token` may look like a JWT... you should treat it as a string and not try to decode it... please. We don't guarentee it'll always be a JWT and want to change it in the background without updating clients.

### `change_handle`

#### Request

```json
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN",
	"new_handle": "GEM"
}
```

#### Response

```json
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN",
	"handle": "GEM",
	"numeric": "323"
}
```

### `release_handle`

#### Request

```json
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN"
}
```

### `change_pin`

#### Request

```json
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN",
	"current_pin": "014410",
	"new_pin": "111111"
}
```

### `list_events`

#### Request

```json
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN",
	"pagination": {
		"before": null,
		"after": "event_000000C0PogVwB5cpZgpzndpiBp4Y",
		"order": "oldest_first",
		"limit": 50
	}
}
```

#### Response

```json
[
	{
		"id": "event_000000C0PLcdyYdxIBi9tOUYjAldI",
		"type": "create_user",
		"user_id": "user_000000C0PLcdyYdxIBi9tOUYjAldJ",
		"game_id": null,
		"payload": {
			"creation_id": "user_000000C0PLcdyYdxIBi9tOUYjAldJ",
			"handle": "DFL",
			"numeric": "001"
		},
		"created_at": "2020-10-17T14:18:29.383619Z"
	},
	{
		"id": "event_000000C0PMw6WdMpoiEl8ZV2SzO3E",
		"type": "change_pin",
		"user_id": "user_000000C0PLcdyYdxIBi9tOUYjAldJ",
		"game_id": null,
		"payload": {
			"user_id": "user_000000C0PLcdyYdxIBi9tOUYjAldJ"
		},
		"created_at": "2020-10-17T14:30:45.029983Z"
	}
]
```

### `test_game_options`

#### Request

```json
{
	"lines": [
		"district",
		"hammersmith-city",
		"london-overground",
		"dlr"
	],
	"zones": [
		"2",
		"3",
		"5"
	]
}
```

#### Response

```json
{
	"possible_prompts": 41
}
```

### `create_game`

#### Request

```js
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN",
	"difficulty_options": {
		"rounds": 5,
		"include_random_spaces": true,
		"change_letter_order": true,
		"reveal_word_length": false,
	},
	"game_options": {
		"lines": [
			"district",
			"waterloo-city",
			...more
		],
		"zones": [
			"1",
			"4",
			"6",
			...more
		]
	}
}
```

#### Response

```json
{
	"id": "game_000000C0HbJWhvF4jnfFcEulxhsaG",
	"next": {
		"prompt_id": "prompt_000000C0SgBJcmd3VoyUXFmRdW9tr",
		"prompt": "W HTC HP L",
		"length": 11
	}
}
```

`length` is `null` if the `reveal_word_length` option is set to `false`.

### `submit_answer`

#### Request

```json
{
	"prompt_id": "prompt_000000C0HbJWhvF4jnfFcEulxhsaG",
	"answer": "whitechapel"
}
```

#### Response

```json
{
	"correct": true,
	"answer": "WHITECHAPEL",
	"next": {
		"prompt_id": "prompt_000000C0HbJWhvF4jnfFcEulxhsaG",
		"prompt": "BR KNG"
	}
}
```

`prompt` will be `null` if the round count has matched. Use `get_game_state` to see the final scores.

### `get_current_game`

#### Request

```json
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN"
}
```

#### Response

```json
{
	"game_id": "game_000000C0HbJWhvF4jnfFcEulxhsaG",
	"next": {
		"prompt_id": "prompt_000000C0SgBJcmd3VoyUXFmRdW9tp",
		"prompt": "VL",
		"length": 4
	}
}
```

### `get_hint`

#### Request

```json
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN",
	"prompt_id": "prompt_000000C0SgBJcmd3VoyUXFmRdW9tp"
}
```

#### Response

```json
{
	"lines": ["central", "waterloo-city"]
}
```

### `get_game_state`

#### Request

```json
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN",
	"game_id": "game_000000C0HbJWhvF4jnfFcEulxhsaG"
}
```

#### Response

```json
{
	"in_progress": false,
	"score": 18,
	"game_time": "P14M25S",
	"difficulty_options": {
		"rounds": 5,
		"include_random_spaces": true,
		"change_letter_order": true,
		"reveal_word_length": false
	},
	"game_options": {
		"lines": [
			"district",
			"bakerloo",
			"circle"
		],
		"bus_stops": false,
		"overground": true
	}
}
```

### `explain_score`

#### Request

```json
{
	"user_id": "user_000000C0KhWzlS5SMpISfDx5IF3aN",
	"game_id": "game_000000C0HbJWhvF4jnfFcEulxhsaG"
}
```

`game_id` is nullable, if so the user score is provided.

#### Response

```json
{
	"start": 0,
	"base": 2,
	"end": 2,
	"final": 2,
	"events": [
		{
			"prompt_id": "prompt_000000C0SgBJcmd3VoyUXFmRdW9tp",
			"effect": "+2",
			"score": 2,
			"note": "Answered & correct"
		}
	]
}
```

### `get_leaderboard`

#### Response

```json
{
	"level": "Bronze",
	"color": "#002002",
	"players": [
		{
			"user_id": "user_000000C0hr6WoXiCOAUWI9NJrF0E9",
			"handle": "DFL",
			"numeric": "001",
			"score": 52
		}
	]
}
```

### `get_game_history`

#### Request

```json
{
	"user_id": "user_000000C0hsG8dDlvxfRVKxjtbUwaY",
}
```

#### Response

```json
[
	{
		"game_id": "game_000000C0h8airLlRVMBiATkuj2c3I",
		"score": 23
	},
	{
		"game_id": "game_000000C0h8airLlRVMBiATkuj2c3K",
		"score": 12
	}
]
```
