const users = [];

function addUser({ id, username, room }) {
  let condition = users.find((item) => item.username === username);

  if (condition) {
    return { error: "this username already used!", user: null };
  }

  users.push({ id, username, room });
  return { user: { id, username, room }, error: null };
}

function getUser(username) {
  const find = users.find((item) => item.username === username);

  if (!find) {
    return { error: "this user not found!", user: null };
  }

  return { user: find, error: null };
}

function removeUser(username) {
  const findIndex = users.findIndex((item) => item.username === username);

  if (findIndex !== -1) {
    users.splice(findIndex, 1);
  }
}

function getUsersInRoom(room) {
  const filterUsers = users.filter((item) => item.room === room);

  if (filterUsers.length <= 0) {
    return { error: "not found!", length: null };
  }

  return { error: null, length: filterUsers.length };
}

function logger() {
  return console.log(users);
}

module.exports = { logger, addUser, getUser, removeUser, getUsersInRoom };
