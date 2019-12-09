### 使用 entity-component-system 架构

### 主逻辑循环 
`
1. update all object by force generator
2. simulate all object's movement or rotate by speed
3. solve contact
4. solve collision
5. update all object velocity, position, rotation 
6. renderer: show it in the graphics-world
`

### 主要数据结构定义：
1. World 世界，掌管所有的物体
2. Context Entity的集合 用于给System做逻辑
3. Entity 实体，持有所有的component 继承自Object
4. Object 只是含有InstanceId 和world的指针
5. Settings 世界内的设置
6. Particle 质点物理的刚体
7. collision 碰撞系统
8. contact ？？