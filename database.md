# Database

This document lays out the database information for the game and it's format

## Events - `events`

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

## Users - `proj_users`

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
			<td>tag</td>
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
			<td>false</td>
			<td>hashed</td>
		</tr>
		<tr>
			<td>score</td>
			<td>int</td>
			<td>false</td>
			<td></td>
		</tr>
	</tbody>
</table>

## Games - `proj_games`

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

## Prompts - `proj_prompts`

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
	</tbody>
</table>
