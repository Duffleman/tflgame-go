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
		"zones": [
			"1",
			"2",
			"5"
		]
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
	"answer_given": "WHITECHAPEL",
	"correct": true,
	"answered_at": "2020-01-01T00:00:00Z"
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

## `change_handle`

```json
{
	"user_id": "user_000000C0HjBIFIIxIKLTiSVviqtpE",
	"new_handle": "GEM",
	"new_numeric": "001"
}
```

## `change_pin`

```json
{
	"user_id": "user_000000C0HjBIFIIxIKLTiSVviqtpE",
	"pin": "<hash>"
}
```

## `release_handle`

```json
{
	"user_id": "user_000000C0HjBIFIIxIKLTiSVviqtpE",
	"handle": "DFL",
	"numeric": "001"
}
```

## `finish_game`

```json
{
	"user_id": "user_000000C0HjBIFIIxIKLTiSVviqtpE",
	"game_id": "game_000000C0HbJWhvF4jnfFcEulxhsaH"
}
```

## `give_hint`

```json
{
	"prompt_id": "prompt_000000C0SgBJcmd3VoyUXFmRdW9tq",
	"new_prompt": "FNSBRY PRK",
	"lines": [
		"district"
	]
}
```
