# Database

This document lays out the database information for the game and it's format

## Game data

### Events - `events`

<table>
	<thead>
		<th>name</th>
		<th>type</th>
		<th>nullable</th>
		<th>notes</th>
	</thead>
	<tbody>
		<tr>
			<td>id</td>
			<td>text</td>
			<td>false</td>
			<td>primary key</td>
		</tr>
		<tr>
			<td>serial</td>
			<td>bigint</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>type</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>user_id</td>
			<td>text</td>
			<td>false</td>
			<td>foreign key to <code>proj_users</code></td>
		</tr>
		<tr>
			<td>game_id</td>
			<td>text</td>
			<td>true</td>
			<td>foreign key to <code>proj_games</code></td>
		</tr>
		<tr>
			<td>payload</td>
			<td>jsonb</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>created_at</td>
			<td>datetime</td>
			<td>false</td>
			<td></td>
		</tr>
	</tbody>
</table>

### Users - `proj_users`

<table>
	<thead>
		<th>name</th>
		<th>type</th>
		<th>nullable</th>
		<th>notes</th>
	</thead>
	<tbody>
		<tr>
			<td>id</td>
			<td>text</td>
			<td>false</td>
			<td>primary key</td>
		</tr>
		<tr>
			<td>handle</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>numeric</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>pin</td>
			<td>text</td>
			<td>true</td>
			<td>hashed</td>
		</tr>
		<tr>
			<td>score</td>
			<td>int</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>created_at</td>
			<td>datetime</td>
			<td>false</td>
			<td></td>
		</tr>
	</tbody>
</table>

### Games - `proj_games`

<table>
	<thead>
		<th>name</th>
		<th>type</th>
		<th>nullable</th>
		<th>notes</th>
	</thead>
	<tbody>
		<tr>
			<td>id</td>
			<td>text</td>
			<td>false</td>
			<td>primary key</td>
		</tr>
		<tr>
			<td>user_id</td>
			<td>text</td>
			<td>false</td>
			<td>foreign key to <code>proj_users</code></td>
		</tr>
		<tr>
			<td>difficulty_options</td>
			<td>jsonb</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>game_options</td>
			<td>jsonb</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>score</td>
			<td>int</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>created_at</td>
			<td>datetime</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>finished_at</td>
			<td>datetime</td>
			<td>true</td>
			<td></td>
		</tr>
	</tbody>
</table>

### Prompts - `proj_prompts`

<table>
	<thead>
		<th>name</th>
		<th>type</th>
		<th>nullable</th>
		<th>notes</th>
	</thead>
	<tbody>
		<tr>
			<td>id</td>
			<td>text</td>
			<td>false</td>
			<td>primary key</td>
		</tr>
		<tr>
			<td>user_id</td>
			<td>text</td>
			<td>false</td>
			<td>foreign key to <code>proj_users</code></td>
		</tr>
		<tr>
			<td>game_id</td>
			<td>text</td>
			<td>false</td>
			<td>foreign key to <code>proj_games</code></td>
		</tr>
		<tr>
			<td>prompt</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>answer</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>answer_given</td>
			<td>text</td>
			<td>true</td>
			<td></td>
		</tr>
		<tr>
			<td>correct</td>
			<td>boolean</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>created_at</td>
			<td>datetime</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>answered_at</td>
			<td>datetime</td>
			<td>true</td>
			<td></td>
		</tr>
		<tr>
			<td>hint_given_at</td>
			<td>datetime</td>
			<td>true</td>
			<td></td>
		</tr>
	</tbody>
</table>

## TFL Data

### Modes - `tfl_modes`

<table>
	<thead>
		<th>name</th>
		<th>type</th>
		<th>nullable</th>
		<th>notes</th>
	</thead>
	<tbody>
		<tr>
			<td>id</td>
			<td>text</td>
			<td>false</td>
			<td>primary key</td>
		</tr>
		<tr>
			<td>name</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
	</tbody>
</table>

### Lines - `tfl_lines`

<table>
	<thead>
		<th>name</th>
		<th>type</th>
		<th>nullable</th>
		<th>notes</th>
	</thead>
	<tbody>
		<tr>
			<td>id</td>
			<td>text</td>
			<td>false</td>
			<td>primary key</td>
		</tr>
		<tr>
			<td>name</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>mode_name</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>created_at</td>
			<td>datetime</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>modified_at</td>
			<td>datetime</td>
			<td>false</td>
			<td></td>
		</tr>
	</tbody>
</table>

### Stops - `tfl_stops`

<table>
	<thead>
		<th>name</th>
		<th>type</th>
		<th>nullable</th>
		<th>notes</th>
	</thead>
	<tbody>
		<tr>
			<td>id</td>
			<td>text</td>
			<td>false</td>
			<td>primary key</td>
		</tr>
		<tr>
			<td>name</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>ics_code</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>station_naptan</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>status</td>
			<td>boolean</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>lat</td>
			<td>float</td>
			<td>false</td>
			<td></td>
		</tr>
		<tr>
			<td>lon</td>
			<td>float</td>
			<td>false</td>
			<td></td>
		</tr>
	</tbody>
</table>

### Lines stops (pivot) - `tfl_lines_stops`

<table>
	<thead>
		<th>name</th>
		<th>type</th>
		<th>nullable</th>
		<th>notes</th>
	</thead>
	<tbody>
		<tr>
			<td>line_id</td>
			<td>text</td>
			<td>false</td>
			<td>primary key</td>
		</tr>
		<tr>
			<td>stop_id</td>
			<td>text</td>
			<td>false</td>
			<td>primary key</td>
		</tr>
		<tr>
			<td>mode</td>
			<td>text</td>
			<td>false</td>
			<td></td>
		</tr>
	</tbody>
</table>
