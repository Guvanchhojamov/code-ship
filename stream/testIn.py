import requests

invidious_instance = "https://some.invidious.instance"
video_id = "dQw4w9WgXcQ"  # Replace with a real video ID

url = f"{invidious_instance}/api/v1/videos/{video_id}"

try:
    response = requests.get(url)
    if response.status_code == 200:
        video_data = response.json()
        # Now you'd need to examine video_data to see if it contains
        # live stream information and how it's structured.
        print(video_data)  # Print the data to inspect it
    else:
        print(f"Error: {response.status_code}")
except requests.exceptions.RequestException as e:
    print(f"Error: {e}")