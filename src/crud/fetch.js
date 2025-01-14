const baseUrl = 'http://your-api-endpoint.com/'; 

async function get(endpoint) {
  try {
    const response = await fetch(`${baseUrl}${endpoint}`);
    if (!response.ok) {
      throw new Error(`Network response was not ok (${response.status})`);
    }
    return await response.json();
  } catch (error) {
    console.error('Error fetching data:', error);
    throw error; // Re-throw the error for proper handling
  }
}

async function post(endpoint, data) {
  try {
    const response = await fetch(`${baseUrl}${endpoint}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
    if (!response.ok) {
      throw new Error(`Network response was not ok (${response.status})`);
    }
    return await response.json();
  } catch (error) {
    console.error('Error creating data:', error);
    throw error;
  }
}

async function put(endpoint, data) {
  try {
    const response = await fetch(`${baseUrl}${endpoint}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
    if (!response.ok) {
      throw new Error(`Network response was not ok (${response.status})`);
    }
    return await response.json();
  } catch (error) {
    console.error('Error updating data:', error);
    throw error;
  }
}

async function del(endpoint) {
  try {
    const response = await fetch(`${baseUrl}${endpoint}`, {
      method: 'DELETE',
    });
    if (!response.ok) {
      throw new Error(`Network response was not ok (${response.status})`);
    }
    return response.status === 204; // Check for successful deletion (no content)
  } catch (error) {
    console.error('Error deleting data:', error);
    throw error;
  }
}

// Example Usage
async function fetchData() {
  try {
    const data = await get('/users'); 
    console.log(data);

    const newData = { name: 'New User', email: 'newuser@example.com' };
    const createdData = await post('/users', newData);
    console.log(createdData);

    const updatedData = { id: createdData.id, name: 'Updated User' };
    const updatedResult = await put(`/users/${updatedData.id}`, updatedData);
    console.log(updatedResult);

    const deleted = await del(`/users/${updatedData.id}`);
    console.log('Deleted:', deleted); 

  } catch (error) {
    console.error('An error occurred:', error);
  }
}

  