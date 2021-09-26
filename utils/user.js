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

function removeUser(id) {
  const findIndex = users.findIndex((item) => item.id === id);

  if (findIndex !== -1) {
    users.splice(findIndex, 1);
  }
}

function getUsersInRoom(room) {
  const filterUsers = users.filter((item) => item.room === room);

  if (filterUsers.length <= 0) {
    return { error: "not found!" };
  }

  return filterUsers;
}

module.exports = { addUser, getUser, removeUser, getUsersInRoom };
