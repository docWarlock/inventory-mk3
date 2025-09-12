import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import apiClient from '../utils/apiClient';

const RoomForm = () => {
    const [room, setRoom] = useState({
        name: '',
        area: 0,
        unit: ''
    });
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const navigate = useNavigate();
    const { id, houseId } = useParams(); // Get both room ID and house ID from route params

    useEffect(() => {
        if (id) {
            fetchRoom(id);
        }
    }, [id]);

    const fetchRoom = async (roomId) => {
        try {
            const response = await apiClient.get(`/rooms/${roomId}`);
            // Ensure we have a valid room object
            setRoom(response.data || { name: '', area: 0, unit: '' });
        } catch (err) {
            setError('Failed to fetch room');
            // Set default empty room state on error
            setRoom({ name: '', area: 0, unit: '' });
        }
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setRoom(prev => ({
            ...prev,
            [name]: value
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError(null);

        try {
            console.log('Submitting room:', room); // Debug log

            // If we're creating a new room and have a houseId, add it to the room data
            let roomData = { ...room };
            if (!id && houseId) {
                roomData.house_id = houseId;
            }

            if (id) {
                // Update existing room
                console.log('Updating room with ID:', id);
                await apiClient.put(`/rooms/${id}`, roomData);
                console.log('Update successful, navigating back');
                // Navigate back to house rooms
                navigate(`/houses/${houseId}/rooms`);
            } else {
                // Create new room - need to check if we're creating for a specific house
                console.log('Creating new room with data:', roomData);
                const response = await apiClient.post('/rooms', roomData);
                console.log('Created room response:', response); // Debug log

                if (houseId) {
                    // Navigate back to house rooms view
                    navigate(`/houses/${houseId}/rooms`);
                } else {
                    // Navigate to all rooms view
                    navigate('/rooms');
                }
            }
        } catch (err) {
            console.error('Error saving room:', err); // Debug log
            console.error('Error details:', err.response?.data || err.message); // More detailed error logging
            setError('Failed to save room: ' + (err.response?.data?.message || err.message || 'Unknown error'));
            setLoading(false);
        }
    };

    return (
        <div>
            <h2>{id ? 'Edit Room' : 'Add New Room'}</h2>
            {error && <div className="error">{error}</div>}
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Name:</label>
                    <input
                        type="text"
                        name="name"
                        value={room?.name || ''}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Area:</label>
                    <input
                        type="number"
                        name="area"
                        value={room?.area || ''}
                        onChange={handleChange}
                    />
                </div>
                <div>
                    <label>Unit:</label>
                    <select
                        name="unit"
                        value={room?.unit || ''}
                        onChange={handleChange}
                    >
                        <option value="">Select unit</option>
                        <option value="sqft">Square Feet</option>
                        <option value="sqm">Square Meters</option>
                        <option value="acres">Acres</option>
                        <option value="hectares">Hectares</option>
                    </select>
                </div>
                <button type="submit" disabled={loading}>
                    {loading ? 'Saving...' : 'Save Room'}
                </button>
            </form>
        </div>
    );
};

export default RoomForm;
