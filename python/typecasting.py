# typecasting str(), int(), float(), bool()

name = "Bro"
age= 24 
gpa = 3.5 
is_student = False 

print(type(name))
print(type(age))
print(type(gpa))
print(type(is_student))

# casting 
gpa = int(gpa)
print(gpa)
age = float(age) 
print(age) 
age = str(age) 
print(age, type(age))

name = ""

print("name: ", name , bool(name))