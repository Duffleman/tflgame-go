# Events

## `create_game`

```json
{
	"user_id": "user_000000C0MOMLVpVxWLULYZphFtGeG",
	"creation_id": "game_000000C0HbJWhvF4jnfFcEulxhsaH",
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

## `submit_prompt`

```json
{
	"game_id": "game_000000C0HbJWhvF4jnfFcEulxhsaH",
	"user_id": "user_000000C0HjBIFIIxIKLTiSVviqtpE",
	"prompt": "W HTC HPL"
}
```

## `answer_prompt`

```json
{
	"game_id": "game_000000C0HbJWhvF4jnfFcEulxhsaH",
	"user_id": "user_000000C0HjBIFIIxIKLTiSVviqtpE",
	"answer": "WHITECHAPEL",
	"correct": true
}
```

## `create_user`

```json
{
	"creation_id": "user_000000C0HjBIFIIxIKLTiSVviqtpE",
	"handle": "DFL",
	"numeric": "001",
	"pin": "<hash>"
}
```

## `change_user_tag`

```json
{
	"user_id": "user_000000C0HjBIFIIxIKLTiSVviqtpE",
	"new_tag": "GEM",
	"new_numeric": "001"
}
```

## `change_user_pin`

```json
{
	"tag": "DFL",
	"numeric": "001",
	"new_pin": "<hash>"
}
```

## `release_user_tag`

```json
{
	"tag": "DFL",
	"numeric": "001"
}
```
